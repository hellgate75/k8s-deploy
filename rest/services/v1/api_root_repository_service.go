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

func getRestV1RepositoryRootApiReference(method string) model.ApiReference {
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

type RestV1RepositoryRootResponse struct {
	Repositories []string `yaml:"repositories,omitempty" json:"repositories,omitempty" xml:"k8srepo,omitempty"`
}

type RestV1RepositoryRootRequest struct {
	Name       string           `yaml:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Id         string           `yaml:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Repository model.Repository `yaml:"repository,omitempty" json:"repository,omitempty" xml:"repository,omitempty"`
}

// RestV1RepositoryRootService is an implementation of RestService interface.
type RestV1RepositoryRootService struct {
	Log                      log.Logger
	BaseUrl                  string
	Configuration            model.KubeRepoConfig
	DataManager              model.RepositoryDataManager
	RepositoryStorageManager model.RepositoryStorageManager
}

// Create is HTTP handler of POST model.Request.
// Use for adding new record to DNS server.
func (s *RestV1RepositoryRootService) Create(w http.ResponseWriter, r *http.Request) {
	s.Log.Infof("RestV1RepositoryRootService.Create() - Path: %s ...", r.URL.Path)
	var request = RestV1RepositoryRootRequest{}
	var response model.Response
	err := utils.RestParseRequest(w, r, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = model.Response{
			Status:    http.StatusBadRequest,
			Message:   fmt.Sprintf("Error parsing request: %v", err),
			Reference: getRestV1RepositoryRootApiReference("POST"),
			Data:      nil,
		}
	} else {
		if request.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			response = model.Response{
				Status:    http.StatusBadRequest,
				Message:   "Repository Name field must be valid and not empty",
				Reference: getRestV1RepositoryRootApiReference("POST"),
				Data:      nil,
			}
		} else {
			var resp = s.DataManager.AddRepository(request.Name)
			if resp.Success {
				w.WriteHeader(http.StatusOK)
				response = model.Response{
					Status:    http.StatusOK,
					Message:   resp.Message,
					Reference: getRestV1RepositoryRootApiReference("POST"),
					Data:      resp.ResponseObjects[0],
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				response = model.Response{
					Status:    http.StatusInternalServerError,
					Message:   fmt.Sprintf("Error creating repository: %s, message: %s", request.Name, resp.Message),
					Reference: getRestV1RepositoryRootApiReference("POST"),
					Data:      nil,
				}
			}
		}
	}
	err = utils.RestParseResponse(w, r, &response)
	if err != nil {
		s.Log.Errorf("Error encoding response: %v", err)
	}
}

// Read is HTTP handler of GET model.Request.
// Use for reading existed records on DNS server.
func (s *RestV1RepositoryRootService) Read(w http.ResponseWriter, r *http.Request) {
	s.Log.Infof("RestV1RepositoryRootService.Read() - Path: %s ...", r.URL.Path)
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
		if method == "" || strings.ToLower(method) == "post" {
			templates = append(templates, rest.TemplateDataType{
				Method:  "POST",
				Header:  []string{},
				Query:   []string{},
				Request: RestV1RepositoryRootRequest{},
			})
		}
		if method == "" || strings.ToLower(method) == "put" {
			templates = append(templates, rest.TemplateDataType{
				Method:  "PUT",
				Header:  []string{},
				Query:   []string{},
				Request: RestV1RepositoryRootRequest{},
			})
		}
		if method == "" || strings.ToLower(method) == "delete" {
			templates = append(templates, rest.TemplateDataType{
				Method:  "DELETE",
				Header:  []string{"EXCLUDING: true|false", "PURGE: true|false"},
				Query:   []string{"excluding=true|false", "purge=true|false"},
				Request: RestV1RepositoryRootRequest{},
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
	response := model.Response{
		Status:    http.StatusOK,
		Message:   message,
		Reference: getRestV1RepositoryRootApiReference("GET"),
		Data:      RestV1RepositoryRootResponse{Repositories: list},
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
func (s *RestV1RepositoryRootService) Update(w http.ResponseWriter, r *http.Request) {
	s.Log.Infof("RestV1RepositoryRootService.Update() - Path: %s ...", r.URL.Path)
	var request = RestV1RepositoryRootRequest{}
	var response model.Response
	err := utils.RestParseRequest(w, r, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = model.Response{
			Status:    http.StatusBadRequest,
			Message:   fmt.Sprintf("Error parsing request: %v", err),
			Reference: getRestV1RepositoryRootApiReference("PUT"),
			Data:      nil,
		}
	} else {
		if request.Id == "" || request.Repository.Name == "" || request.Repository.Id == "" {
			w.WriteHeader(http.StatusBadRequest)
			response = model.Response{
				Status:    http.StatusBadRequest,
				Message:   "Repository Id and Repository Body field must be valid and not empty",
				Reference: getRestV1RepositoryRootApiReference("PUT"),
				Data:      nil,
			}
		} else {
			var resp = s.DataManager.UpdateRepository(request.Id, &request.Repository)
			if resp.Success {
				w.WriteHeader(http.StatusOK)
				response = model.Response{
					Status:    http.StatusOK,
					Message:   resp.Message,
					Reference: getRestV1RepositoryRootApiReference("PUT"),
					Data:      resp.ResponseObjects[0],
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				response = model.Response{
					Status:    http.StatusInternalServerError,
					Message:   fmt.Sprintf("Error updating repository: %s, message: %s", request.Name, resp.Message),
					Reference: getRestV1RepositoryRootApiReference("PUT"),
					Data:      nil,
				}
			}
		}
	}
	err = utils.RestParseResponse(w, r, &response)
	if err != nil {
		s.Log.Errorf("Error encoding response: %v", err)
	}
}

// Delete is HTTP handler of DELETE model.Request.
// Use for removing records on DNS server.
func (s *RestV1RepositoryRootService) Delete(w http.ResponseWriter, r *http.Request) {
	s.Log.Infof("RestV1RepositoryRootService.Delete() - Path: %s ...", r.URL.Path)
	var request = RestV1RepositoryRootRequest{}
	var response model.Response
	err := utils.RestParseRequest(w, r, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = model.Response{
			Status:    http.StatusBadRequest,
			Message:   fmt.Sprintf("Error parsing request: %v", err),
			Reference: getRestV1RepositoryRootApiReference("DELETE"),
			Data:      nil,
		}
	} else {
		if request.Name == "" && request.Id == "" {
			w.WriteHeader(http.StatusBadRequest)
			response = model.Response{
				Status:    http.StatusBadRequest,
				Message:   "Repository Name and/or Repository Id field must be valid and not empty",
				Reference: getRestV1RepositoryRootApiReference("DELETE"),
				Data:      nil,
			}
		} else {
			var items = make([]model.QueryItem, 0)
			var qId = strings.TrimSpace(request.Id)
			if qId != "" {
				items = append(items, model.QueryItem{
					Key:        "id",
					Aggregator: model.AggregatorEq,
					Value:      qId,
				})
			}
			var qName = strings.TrimSpace(request.Name)
			if qName != "" {
				items = append(items, model.QueryItem{
					Key:        "name",
					Aggregator: model.AggregatorEq,
					Value:      qName,
				})
			}
			var q = model.Query{
				Items: items,
				Oper:  model.OperOr,
			}
			var inclusive = true
			if s := r.URL.Query().Get("excluding"); s != "" {
				if parseBool(s) {
					inclusive = false
				}
			}
			if s := r.Header.Get("EXCLUDING"); s != "" {
				if parseBool(s) {
					inclusive = false
				}
			}
			var purge = true
			if s := r.URL.Query().Get("purge"); s != "" {
				if parseBool(s) {
					purge = false
				}
			}
			if s := r.Header.Get("PURGE"); s != "" {
				if parseBool(s) {
					purge = false
				}
			}
			var resp = s.DataManager.DeleteRepositories(inclusive, q)
			if resp.Success && purge {
				s.Log.Warnf("Purging repositories ....")
				resp = s.DataManager.PurgeRepositories(inclusive, q)
			}
			if resp.Success {
				w.WriteHeader(http.StatusOK)
				response = model.Response{
					Status:    http.StatusOK,
					Message:   resp.Message,
					Reference: getRestV1RepositoryRootApiReference("DELETE"),
					Data:      resp.ResponseObjects,
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				response = model.Response{
					Status:    http.StatusInternalServerError,
					Message:   fmt.Sprintf("Error deleting repository-> name: <%s>, id: <%s>, message: %s", request.Id, request.Name, resp.Message),
					Reference: getRestV1RepositoryRootApiReference("DELETE"),
					Data:      nil,
				}
			}
		}
	}
	err = utils.RestParseResponse(w, r, &response)
	if err != nil {
		s.Log.Errorf("Error encoding response: %v", err)
	}
}
func parseBool(s string) bool {
	var sp = strings.ToLower(strings.TrimSpace(s))
	return sp == "true" || sp == "1" || sp == "yes"
}
