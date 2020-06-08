package services

import (
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/rest/services/v1"
	"net/http"
)

// Interface that describes the main feature of a Rest Service, used by a Rest Server
type RestService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// Creates a V1 API Rest Service Instance
func NewV1RegistryRootRestService(logger log.Logger, hostBaseUrl string,
	configuration model.KubeRepoConfig,
	dataManager model.RepositoryDataManager,
	repositoryStorageManager model.RepositoryStorageManager) RestService {
	return &v1.RestV1RepositoryRootService{
		Log:                      logger,
		BaseUrl:                  hostBaseUrl,
		Configuration:            configuration,
		DataManager:              dataManager,
		RepositoryStorageManager: repositoryStorageManager,
	}
}
