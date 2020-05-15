// Copyright 2020 Re-Bind Author (Fabrizio Torelli). All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hellgate75/k8s-deploy/integration"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/model/rest"
	"github.com/hellgate75/k8s-deploy/rest/services"
	"github.com/hellgate75/k8s-deploy/utils"
	"net/http"
	"os"
)

var rwDirPath string
var configDirPath string
var initializeAndExit bool
var useConfigFile bool
var enableFileLogging bool
var logVerbosity string
var logFilePath string
var mongoDbEnabled bool
var mongoDbHost string
var mongoDbPort int
var mongoDbUser string
var mongoDbPassword string
var enableLogRotate bool
var logMaxFileSize int64
var logMaxFileCount int
var listenIP string
var listenPort int
var tlsCert string
var tlsKey string

//TODO: Give Life to Logger
var logger log.Logger = log.NewLogger("k8s-deploy", log.DEBUG)

func init() {
	integration.InitPackage()
	logger.Info("Initializing K8s Repo Rest Server ....")
	flag.StringVar(&rwDirPath, "data-dir", rest.DefaultStorageFolder, "dns storage dir")
	flag.StringVar(&configDirPath, "config-dir", rest.DefaultConfigFolder, "dns config dir")
	flag.BoolVar(&initializeAndExit, "init-and-exit", false, "initialize config in the config dir and exit")
	flag.BoolVar(&useConfigFile, "use-config-file", false, "use config file instead parameters")
	flag.BoolVar(&enableFileLogging, "enable-file-log", false, "enable logginf over file")
	flag.StringVar(&logVerbosity, "log-verbosity", rest.DefaultLogFileLevel, "log file verbosity level (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)")
	flag.StringVar(&logFilePath, "log-file-path", rest.DefaultLogFileFolder, "log file path")
	flag.BoolVar(&enableLogRotate, "log-rotate", true, "log file rotation enabled")
	flag.Int64Var(&logMaxFileSize, "log-max-size", 1024, "log file rotation max file size in bytes")
	flag.IntVar(&logMaxFileCount, "log-count", 1024, "log file rotation max number of file")
	flag.StringVar(&listenIP, "listen-ip", rest.DefaultIpAddress, "http server ip")
	flag.IntVar(&listenPort, "listen-port", rest.DefaultRestServerPort, "http server port")
	flag.StringVar(&tlsCert, "tsl-cert", "", "tls certificate file path")
	flag.StringVar(&tlsKey, "tsl-key", "", "tls certificate key file path")
	flag.BoolVar(&mongoDbEnabled, "mongo-db-enabled", false, "MongoDb enabled state")
	flag.StringVar(&mongoDbHost, "mongo-db-host", "127.0.0.1", "MongoDb hostname or public ip")
	flag.IntVar(&mongoDbPort,"mongo-db-host", 27017, "MongoDb public port")
	flag.StringVar(&mongoDbUser, "mongo-db-user", "", "MongoDb user name")
	flag.StringVar(&mongoDbPassword, "mongo-db-password", "", "MongoDb user password")
}

func main() {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
			fmt.Sprintf("Error during start-up: %s", err.Error())
		}
	}()
	flag.Parse()
	if utils.StringsListContainItem("-h", flag.Args(), true) ||
		utils.StringsListContainItem("--help", flag.Args(), true) {
		flag.Usage()
		os.Exit(0)
	}
	if initializeAndExit {
		logger.Info("Initialize K8s Repo Rest Server and Exit!!")
		config := model.KubeRepoConfig{
			DataDirPath:         rwDirPath,
			ConfigDirPath:       configDirPath,
			ListenIP:            listenIP,
			ListenPort:          listenPort,
			TlsCert:             tlsCert,
			TlsKey:              tlsKey,
			EnableFileLogging:   enableFileLogging,
			LogVerbosity:        logVerbosity,
			LogFilePath:         logFilePath,
			LogFileCount:        logMaxFileCount,
			LogMaxFileSize:      logMaxFileSize,
			EnableLogRotate:     enableLogRotate,
			MongoDbEnabled: 	 mongoDbEnabled,
			MongoDbHost:		 mongoDbHost,
			MongoDbPort:		 mongoDbPort,
			MongoDbUser:		 mongoDbUser,
			MongoDbPassword:	 mongoDbPassword,
		}
		cSErr := model.SaveConfig(configDirPath, "reweb", &config)
		if cSErr != nil {
			logger.Errorf("Unable to save default config to file: ", cSErr)
		}
		os.Exit(0)
	}
	if useConfigFile {
		logger.Warn("Initialize K8s Repo from config file ...")
		logger.Warnf("K8s Repo config folder: %s", configDirPath)
		var config model.KubeRepoConfig
		cLErr := model.LoadConfig(configDirPath, "reweb", &config)
		if cLErr != nil {
			logger.Errorf("Unable to load default config from file: ", cLErr)
		} else {
			logger.Warn("Loading configuration from file complete!!")
			logger.Debugf("Configuration: %s", config.ToJson())
			rwDirPath = config.DataDirPath
			configDirPath = config.ConfigDirPath
			listenIP = config.ListenIP
			listenPort = config.ListenPort
			enableFileLogging = config.EnableFileLogging
			logVerbosity = config.LogVerbosity
			logFilePath = config.LogFilePath
			logMaxFileCount = config.LogFileCount
			logMaxFileSize = config.LogMaxFileSize
			enableLogRotate = config.EnableLogRotate
			tlsCert = config.TlsCert
			tlsKey = config.TlsKey
			mongoDbEnabled = config.MongoDbEnabled
			mongoDbHost = config.MongoDbHost
			mongoDbPort = config.MongoDbPort
			mongoDbUser = config.MongoDbUser
			mongoDbPassword = config.MongoDbPassword
		}
	}
	verbosity := log.LogLevelFromString(logVerbosity)
	logger.Warnf("File logging enabled: %v", enableFileLogging)
	if enableFileLogging {
		logger.Warnf("Enabling file logging at path: %s", logFilePath)
		if _, err := os.Stat(logFilePath); err != nil {
			_ = os.MkdirAll(logFilePath, 0660)
		}
		logger.Warnf("File logging at path: %s enabled...", logFilePath)
		logDir, _ := os.Open(logFilePath)
		var logErr, logRErr error
		var rotator log.LogRotator
		if enableLogRotate {
			logger.Warn("Enable log rotating ...")
			rotator, logRErr = log.NewLogRotator(logDir, "reweb.log", logMaxFileSize, logMaxFileCount, nil)
		} else {
			logger.Warn("No log rotating is enabled...")
			rotator, logRErr = log.NewLogNoRotator(logDir, "reweb.log", nil)
		}
		if logRErr != nil {
			logger.Errorf("Unable to instantiate log rotator: ", logRErr)
		} else {
			logger.Warn("Starting file logging ...")
			logger, logErr = log.NewFileLogger("re-web",
				rotator,
				verbosity)
			if logErr != nil {
				logger.Warn("No File logging started for error...")
				logger = log.NewLogger("re-web", verbosity)
				logger.Errorf("Unable to instantiate file logger: ", logErr)
			} else {
				logger.Warn("File logging started!!")
			}
		}
	} else {
		logger.Warn("No File logging selected ...")
		logger = log.NewLogger("re-web", verbosity)
	}
	logger.Info("Starting K8s Repo Rest Server ...")
	if err := os.MkdirAll(rwDirPath, 0666); err != nil {
		logger.Errorf("Create rwdirpath: %v error: %v", rwDirPath, err)
		return
	}

	// Create Data Store
	// TODO Read from remote db or local folder
//	store := registry.NewStore(logger, rwDirPath, defaultForwarders, rest.DefaultDomains)
//	store.Load()

	// Handler stuf for the API service groups
	apiHandler := func(service services.RestService) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				service.Create(w, r)
			case http.MethodGet:
				service.Read(w, r)
			case http.MethodPut:
				service.Update(w, r)
			case http.MethodDelete:
				service.Delete(w, r)
			}
		}
	}

	withAuth := func(h http.HandlerFunc) http.HandlerFunc {
		// authentication intercepting
		var _ = "intercept"
		return func(w http.ResponseWriter, r *http.Request) {
			h(w, r)
		}
	}

	rtr := mux.NewRouter()
	var proto string = "http"
	if tlsCert != "" && tlsKey != "" {
		proto = "https"
	}
	// Creates/Sets API endpoints handlers
	services.CreateApiEndpoints(rtr, withAuth, apiHandler,
		logger, fmt.Sprintf("%s://%s:%v", proto, listenIP, listenPort))

	//Adding entry point for generic queries (GET)
	http.Handle("/", rtr)
	// Adding TLS certificates if required
	if tlsCert == "" || tlsKey == "" {
		logger.Infof("RestService start-up:: Starting server in simple mode on ip: %s and port: %v\n", listenIP, listenPort)
		err = http.ListenAndServe(fmt.Sprintf("%s:%v", listenIP, listenPort), nil)
	} else {
		logger.Infof("RestService start-up:: Starting server in simple mode on ip: %s and port: %v\n", listenIP, listenPort)
		logger.Infof("RestService start-up:: Using certificate file: %s and certticate key file: %v..\n", tlsCert, tlsKey)
		err = http.ListenAndServeTLS(fmt.Sprintf("%s:%v", listenIP, listenPort), tlsCert, tlsKey, nil)
	}
	if err != nil {
		logger.Fatalf("RestService start-up:: Error listening on s:%v - Error: %v\n", listenIP, listenPort, err)
		os.Exit(1)
	}
	logger.Info("K8s Repo Server started!!")
}
