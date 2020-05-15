package services

import (
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/rest/services/v1"
	"net/http"
)

type RestService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewV1DnsRootRestService(logger log.Logger, hostBaseUrl string) RestService {
	return &v1.RestRootService{
		Log:     logger,
		BaseUrl: hostBaseUrl,
	}
}
