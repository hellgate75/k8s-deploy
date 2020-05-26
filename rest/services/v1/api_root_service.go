package v1

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/model/rest"
	"github.com/hellgate75/k8s-deploy/utils"
	"net/http"
	"strings"
)

func getRootRepositoryApiReference(method string) model.ApiReference {
	var items = make([]model.ApiReferenceItem, 0)
	items = append(items, model.ApiReferenceItem{
		"GET",
		"/v1/repositories",
	})
	items = append(items, model.ApiReferenceItem{
		"POST",
		"/v1/repositories",
	})
	items = append(items, model.ApiReferenceItem{
		"PUT",
		"/v1/repositories",
	})
	items = append(items, model.ApiReferenceItem{
		"DELETE",
		"/v1/repositories",
	})
	return model.ApiReference{
		CurrentUrl:    "/v1/repositories",
		CurrentMethod: method,
		Urls:          items,
	}
}

type RestRegistryRootResponse struct {
	Repositories []string `yaml:"repositories,omitempty" json:"repositories,omitempty" xml:"k8srepo,omitempty"`
}

// RestRegistryRootService is an implementation of RestService interface.
type RestRegistryRootService struct {
	Log                      log.Logger
	BaseUrl                  string
	Configuration            model.KubeRepoConfig
	DataManager              model.RepositoryDataManager
	RepositoryStorageManager model.RepositoryStorageManager
}

// Create is HTTP handler of POST model.Request.
// Use for adding new record to DNS server.
func (s *RestRegistryRootService) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	response := model.Response{
		Status:    http.StatusMethodNotAllowed,
		Message:   "Not allowed on rest root",
		Reference: getRootRepositoryApiReference("POST"),
		Data:      nil,
	}
	err := utils.RestParseResponse(w, r, &response)
	if err != nil {
		s.Log.Errorf("Error encoding response: %v", err)
	}
}

// Read is HTTP handler of GET model.Request.
// Use for reading existed records on DNS server.
func (s *RestRegistryRootService) Read(w http.ResponseWriter, r *http.Request) {
	var action = r.URL.Query().Get("action")
	var method = r.URL.Query().Get("method")
	if strings.ToLower(action) == "template" {
		var templates = make([]rest.TemplateDataType, 0)
		if method == "" || strings.ToLower(method) == "get" {
			templates = append(templates, rest.TemplateDataType{
				Method:  "GET",
				Header:  []string{},
				Query:   []string{"action=template"},
				Request: nil,
			})
		}
		tErr := utils.RestParseResponse(w, r,
			&rest.TemplateResponse{
				Templates: templates,
			})
		if tErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Log.Errorf("Error encoding template(s) summary response, Error: %v", tErr)
		}
		return
	}
	//	groups := s.Store.GetGroupBucket().ListGroups()
	var list = make([]string, 0)
	var message = "OK"
	resp := s.DataManager.ListRepositories()
	if resp.Success {
		if resp.ResponseObjects != nil {
			for _, obj := range resp.ResponseObjects {
				var r = obj.(model.Repository)
				list = append(list, fmt.Sprintf("%s:%s", r.Id, r.Name))
			}
		}
	} else {
		message = fmt.Sprintf("ERROR:: %s", resp.Message)
	}
	//	for _, g := range groups {
	//		list = append(list, g.Name)
	//	}
	response := model.Response{
		Status:    http.StatusOK,
		Message:   message,
		Reference: getRootRepositoryApiReference("POST"),
		Data:      RestRegistryRootResponse{Repositories: list},
	}
	w.WriteHeader(http.StatusOK)
	err := utils.RestParseResponse(w, r, &response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.Log.Errorf("Error encoding response: %v", err)
	}
}

// Update is HTTP handler of PUT model.Request.
// Use for updating existed records on DNS server.
func (s *RestRegistryRootService) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	response := model.Response{
		Status:  http.StatusMethodNotAllowed,
		Message: "Not allowed on rest root",
		Data:    nil,
	}
	err := utils.RestParseResponse(w, r, &response)
	if err != nil {
		s.Log.Errorf("Error encoding response: %v", err)
	}
}

// Delete is HTTP handler of DELETE model.Request.
// Use for removing records on DNS server.
func (s *RestRegistryRootService) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	response := model.Response{
		Status:  http.StatusMethodNotAllowed,
		Message: "Not allowed on rest root",
		Data:    nil,
	}
	err := utils.RestParseResponse(w, r, &response)
	if err != nil {
		s.Log.Errorf("Error encoding response: %v", err)
	}
}
