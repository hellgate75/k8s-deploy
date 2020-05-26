package services

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"net/http"
)

type EndPointType byte

const (
	RepositoryEndpoint EndPointType = iota + 1
)

func CreateApiEndpoints(router *mux.Router,
	authFunc func(h http.HandlerFunc) http.HandlerFunc,
	dnsHandler func(serv RestService) http.HandlerFunc,
	logger log.Logger,
	hostBaseUrl string,
	epType EndPointType,
	configuration interface{},
	dataManager model.RepositoryDataManager,
	repositoryStorageManager model.RepositoryStorageManager) error {
	switch epType {
	case RepositoryEndpoint:
		createV1RegistryApiEndpoints(router, authFunc, dnsHandler, logger, hostBaseUrl, configuration.(model.KubeRepoConfig), dataManager, repositoryStorageManager)
		return nil
	default:
		return errors.New("Not implemented")
	}
	//Create v1 APIs
}

func createV1RegistryApiEndpoints(router *mux.Router,
	authFunc func(h http.HandlerFunc) http.HandlerFunc,
	restHandler func(serv RestService) http.HandlerFunc,
	logger log.Logger,
	hostBaseUrl string,
	config model.KubeRepoConfig,
	dataManager model.RepositoryDataManager,
	repositoryStorageManager model.RepositoryStorageManager) {
	v1RegistryRootRest := NewV1RegistryRootRestService(logger, hostBaseUrl, config, dataManager, repositoryStorageManager)
	router.HandleFunc("/v1/repositories", authFunc(restHandler(v1RegistryRootRest))).Methods("GET", "POST", "PUT", "DELETE")
	//Adding entry point for groups queries (PUT, POST, DEL, GET)
	//router.HandleFunc("/v1/dns/groups", authFunc(restHandler(v1GroupsRest))).Methods("GET", "POST", "PUT", "DELETE")
	////Adding entry point for spcific group queries (PUT, POST, DEL, GET)
	//router.HandleFunc("/v1/dns/group/{group:[a-zA-Z0-9.-]+}", authFunc(restHandler(v1GroupRest))).Methods("GET", "POST", "PUT", "DELETE")
	////Adding entry point for specific group queries (PUT, POST, DEL, GET)
	//router.HandleFunc("/v1/dns/group/{group:[a-zA-Z0-9.-]+}/resources", authFunc(restHandler(v1DnsGroupResourcesRest))).Methods("GET", "POST", "PUT", "DELETE")
	////Adding entry point for specific group queries (PUT, POST, DEL, GET)
	//router.HandleFunc("/v1/dns/group/{group:[a-zA-Z0-9.-]+}/resources/{resource:[a-zA-Z0-9.-]+}", authFunc(restHandler(v1DnsGroupResourceDetailsRest))).Methods("GET", "POST", "PUT", "DELETE")
}
