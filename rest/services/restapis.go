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
	// Kubernetes Repository Api Endpoint constant value
	RepositoryEndpoint EndPointType = iota + 1
	// Kubernetes Deploy Scheduler Api Endpoint constant value
	SchedulerEndpoint
	// Kubernetes Deploy Executor Api Endpoint constant value
	ExecutorEndpoint
)

// Creates the API endpoints, profiled by EndPointType, accordingly to the requesting command.
// Any command will have multiple sections/groups of APIs related to wanted
// feature and service
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
		addV1RepositoryApiEndpoints(router, authFunc, dnsHandler, logger, hostBaseUrl, configuration.(model.KubeRepoConfig), dataManager, repositoryStorageManager)
		return nil
	default:
		return errors.New("Not implemented")
	}
	//Create v1 APIs
}

// Add V1 Repository Endpoints to the Godzilla Router
func addV1RepositoryApiEndpoints(router *mux.Router,
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
