package v1

import (
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/model/rest"
	"github.com/hellgate75/k8s-deploy/utils"
	"net/http"
	"strings"
)

type RestRootResponse struct {
	Repositories []string `yaml:"repositories,omitempty" json:"repositories,omitempty" xml:"repository,omitempty"`
}

// RestRootService is an implementation of RestService interface.
type RestRootService struct {
	Log     log.Logger
	BaseUrl string
}

// Create is HTTP handler of POST model.Request.
// Use for adding new record to DNS server.
func (s *RestRootService) Create(w http.ResponseWriter, r *http.Request) {
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

// Read is HTTP handler of GET model.Request.
// Use for reading existed records on DNS server.
func (s *RestRootService) Read(w http.ResponseWriter, r *http.Request) {
	var action = r.URL.Query().Get("action")
	if strings.ToLower(action) == "template" {
		var templates = make([]rest.TemplateDataType, 0)
		templates = append(templates, rest.TemplateDataType{
			Method:  "GET",
			Header:  []string{},
			Query:   []string{"action=template"},
			Request: nil,
		})
		tErr := utils.RestParseResponse(w, r, &rest.TemplateResponse{
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
//	for _, g := range groups {
//		list = append(list, g.Name)
//	}
	response := model.Response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    RestRootResponse{Repositories: list},
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
func (s *RestRootService) Update(w http.ResponseWriter, r *http.Request) {
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
func (s *RestRootService) Delete(w http.ResponseWriter, r *http.Request) {
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
