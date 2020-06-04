// Copyright 2020 Re-Bind Author (Fabrizio Torelli). All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hellgate75/go-services/database/mongodb"
	"github.com/hellgate75/k8s-deploy/data"
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
var storageNamePrefix string
var enableLogRotate bool
var logMaxFileSize int64
var logMaxFileCount int
var listenIP string
var listenPort int
var tlsCert string
var tlsKey string

const (
	LoggerAppName       = "k8s-deploy-repository"
	ApplicationFullName = "Kubernetes Deploy Repository"
)

//TODO: Give Life to Logger
var logger log.Logger = log.NewLogger(LoggerAppName, log.DEBUG)

func init() {
	integration.InitPackage()
	logger.Infof("Initializing %s Rest Server ....", ApplicationFullName)
	flag.StringVar(&rwDirPath, "data-dir", rest.DefaultRepositoryStorageFolder, "dns storage dir")
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
	flag.IntVar(&listenPort, "listen-port", rest.DefaultRepositoryRestServerPort, "http server port")
	flag.StringVar(&tlsCert, "tsl-cert", "", "tls certificate file path")
	flag.StringVar(&tlsKey, "tsl-key", "", "tls certificate key file path")
	flag.BoolVar(&mongoDbEnabled, "mongo-db-enabled", false, "MongoDb enabled state")
	flag.StringVar(&mongoDbHost, "mongo-db-host", "127.0.0.1", "MongoDb hostname or public ip")
	flag.IntVar(&mongoDbPort, "mongo-db-port", 27017, "MongoDb public port")
	flag.StringVar(&mongoDbUser, "mongo-db-user", "", "MongoDb user name")
	flag.StringVar(&mongoDbPassword, "mongo-db-password", "", "MongoDb user password")
	flag.StringVar(&storageNamePrefix, "storage-name-prefix", rest.DefaultDatabaseNamePrefix, "MongoDb database name or device folder prefix")
}

func main() {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
			fmt.Sprintf("%s:: Error during start-up: %s", ApplicationFullName, err.Error())
		}
	}()
	flag.Parse()
	if utils.StringsListContainItem("-h", flag.Args(), true) ||
		utils.StringsListContainItem("--help", flag.Args(), true) {
		flag.Usage()
		os.Exit(0)
	}
	config := model.KubeRepoConfig{
		DataDirPath:       rwDirPath,
		ConfigDirPath:     configDirPath,
		ListenIP:          listenIP,
		ListenPort:        listenPort,
		TlsCert:           tlsCert,
		TlsKey:            tlsKey,
		EnableFileLogging: enableFileLogging,
		LogVerbosity:      logVerbosity,
		LogFilePath:       logFilePath,
		LogFileCount:      logMaxFileCount,
		LogMaxFileSize:    logMaxFileSize,
		EnableLogRotate:   enableLogRotate,
		MongoDbEnabled:    mongoDbEnabled,
		MongoDbHost:       mongoDbHost,
		MongoDbPort:       mongoDbPort,
		MongoDbUser:       mongoDbUser,
		MongoDbPassword:   mongoDbPassword,
		StorageNamePrefix: storageNamePrefix,
	}
	if initializeAndExit {
		logger.Infof("Initialize %s Rest Server and Exit!!", ApplicationFullName)
		cSErr := model.SaveConfig(configDirPath, LoggerAppName, &config)
		if cSErr != nil {
			logger.Errorf("%s is unable to save default config to file: ", ApplicationFullName, cSErr)
		}
		os.Exit(0)
	}
	if useConfigFile {
		logger.Warnf("Initialize %s from config file ...", ApplicationFullName)
		logger.Warnf("%s config folder: %s", ApplicationFullName, configDirPath)
		var config model.KubeRepoConfig
		cLErr := model.LoadConfig(configDirPath, "k8s-deploy-k8srepo", &config)
		if cLErr != nil {
			logger.Errorf("%s is unable to load default config from file: ", ApplicationFullName, cLErr)
		} else {
			logger.Warnf("%s:: Loading configuration from file complete!!", ApplicationFullName)
			logger.Debugf("%s:: Configuration: %s", ApplicationFullName, config.ToJson())
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
			storageNamePrefix = config.StorageNamePrefix
		}
	}
	verbosity := log.LogLevelFromString(logVerbosity)
	logger.Warnf("%s has file logging enabled: %v", ApplicationFullName, enableFileLogging)
	if enableFileLogging {
		logger.Warnf("%s is enabling file logging at path: %s", ApplicationFullName, logFilePath)
		if _, err := os.Stat(logFilePath); err != nil {
			_ = os.MkdirAll(logFilePath, 0660)
		}
		logger.Warnf("%s uses file logging at path: %s enabled...", ApplicationFullName, logFilePath)
		logDir, _ := os.Open(logFilePath)
		var logErr, logRErr error
		var rotator log.LogRotator
		if enableLogRotate {
			logger.Warnf("%s is enabling log rotating ...", ApplicationFullName)
			rotator, logRErr = log.NewLogRotator(logDir, fmt.Sprintf("rest-%s.log", LoggerAppName), logMaxFileSize, logMaxFileCount, nil)
		} else {
			logger.Warnf("%s has no log rotating is enabled...", ApplicationFullName)
			rotator, logRErr = log.NewLogNoRotator(logDir, fmt.Sprintf("rest-%s.log", LoggerAppName), nil)
		}
		if logRErr != nil {
			logger.Errorf("%s is unable to instantiate log rotator: ", ApplicationFullName, logRErr)
		} else {
			logger.Warnf("%s is starting file logging ...", ApplicationFullName)
			logger, logErr = log.NewFileLogger(LoggerAppName,
				rotator,
				verbosity)
			if logErr != nil {
				logger.Warnf("%s has no File logging started for error...", ApplicationFullName)
				logger = log.NewLogger(LoggerAppName, verbosity)
				logger.Errorf("%s is unable to instantiate file logger: ", ApplicationFullName, logErr)
			} else {
				logger.Warnf("%s:: File logging started!!", ApplicationFullName)
			}
		}
	} else {
		logger.Warnf("%s has no File logging selected ...", ApplicationFullName)
		logger = log.NewLogger(LoggerAppName, verbosity)
	}
	logger.Infof("Starting %s Rest Server ...", ApplicationFullName)
	if err := os.MkdirAll(rwDirPath, 0666); err != nil {
		logger.Errorf("%s:: Create rwdirpath: %v error: %v", ApplicationFullName, rwDirPath, err)
		return
	}

	repositoryStorageManager, err := integration.GetRepositoryStorageManagerSingleton(rwDirPath, logger)
	if err != nil {
		logger.Fatalf("Unable to instantiate Repository storage manager, Error: %s", err.Error())
		os.Exit(1)
	}
	// Create Data Store
	var dataManager model.RepositoryDataManager
	// Read from remote db or local folder
	if mongoDbEnabled {
		driver := mongodb.GetMongoDriver()
		conn, err := driver.Connect(mongodb.GetMongoDbConfig(
			mongoDbHost,
			mongoDbPort,
			mongoDbUser,
			mongoDbPassword))
		if err != nil {
			logger.Fatalf("%s is unable to connect to mongo db, reason: %s", ApplicationFullName, err.Error())
			os.Exit(1)
		}
		dataManager = data.GetMongoRepositoryDataManager(conn, rwDirPath, repositoryStorageManager, logger)
	} else {
		dataManager = data.GetDeviceRepositoryDataManager(rwDirPath, repositoryStorageManager, logger)
	}
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
	err = services.CreateApiEndpoints(rtr, withAuth, apiHandler,
		logger, fmt.Sprintf("%s://%s:%v", proto, listenIP, listenPort),
		services.RepositoryEndpoint, config, dataManager, repositoryStorageManager)
	if err != nil {
		logger.Infof("%s RestService start-up:: Error creating API endpoints: %s\n", ApplicationFullName, err.Error())
		os.Exit(1)
	}
	//Adding entry point for generic queries (GET)
	http.Handle("/", rtr)
	// Adding TLS certificates if required
	if tlsCert == "" || tlsKey == "" {
		logger.Infof("%s RestService start-up:: Starting server in simple mode on ip: %s and port: %v\n", ApplicationFullName, listenIP, listenPort)
		err = http.ListenAndServe(fmt.Sprintf("%s:%v", listenIP, listenPort), nil)
	} else {
		logger.Infof("%s RestService start-up:: Starting server in simple mode on ip: %s and port: %v\n", ApplicationFullName, listenIP, listenPort)
		logger.Infof("%s RestService start-up:: Using certificate file: %s and certticate key file: %v..\n", ApplicationFullName, tlsCert, tlsKey)
		err = http.ListenAndServeTLS(fmt.Sprintf("%s:%v", listenIP, listenPort), tlsCert, tlsKey, nil)
	}
	if err != nil {
		logger.Fatalf("%s RestService start-up:: Error listening on s:%v - Error: %v\n", ApplicationFullName, listenIP, listenPort, err)
		os.Exit(1)
	}
	logger.Infof("%s RestService started!!", ApplicationFullName)
}
