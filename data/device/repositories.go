package device

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
)

type repositoryManager struct {
	baseDataFolder string
	manager        model.RepositoryStorageManager
	logger         log.Logger
}

func (rn *repositoryManager) ListRepositories() model.DataResponse {
	rn.logger.Infof("DeviceRepositoryManager::ListRepositories() ...")
	rl := rn.manager.GetRepositoryList()
	var ril = make([]interface{}, 0)
	for _, r := range rl {
		ril = append(ril, r)
	}
	rn.logger.Infof("DeviceRepositoryManager::ListRepositories() - repos: %v", len(ril))
	return model.DataResponse{
		Success:         true,
		Message:         "OK",
		ResponseObjects: ril,
	}
}

func (rn *repositoryManager) AddRepository(n string) model.DataResponse {
	var repoName = utils.ConvertName(n)
	r, err := rn.manager.CreateRepository(repoName)
	var response = make([]interface{}, 0)
	if err != nil {
		return model.DataResponse{
			Success:         false,
			Message:         fmt.Sprintf("Error creating repository: %s, error: %v", n, err),
			ResponseObjects: response,
		}
	} else {
		response = append(response, *r)
		return model.DataResponse{
			Success:         true,
			Message:         "REPOSITORY CREATED",
			ResponseObjects: response,
		}
	}
}

//Add new k8srepo
func (rn *repositoryManager) UpdateRepository(id string, r *model.Repository) model.DataResponse {
	if r == nil {
		return model.DataResponse{
			Success:         false,
			Message:         fmt.Sprintf("Error updating repository id : %s with empty body", id),
			ResponseObjects: nil,
		}
	}
	rp, err := rn.manager.UpdateRepository(id, *r)
	var response = make([]interface{}, 0)
	if err != nil {
		return model.DataResponse{
			Success:         false,
			Message:         fmt.Sprintf("Error updating repository id: %s, error: %v", id, err),
			ResponseObjects: response,
		}
	} else {
		response = append(response, *rp)
		return model.DataResponse{
			Success:         true,
			Message:         "REPOSITORY Updated",
			ResponseObjects: response,
		}
	}
}

func (rn *repositoryManager) checkFilter(r model.Repository, inclusive bool, q ...model.Query) bool {
	if len(q) == 0 {
		return true
	}
	for _, qr := range q {
		for _, qi := range qr.Items {
			var sKey = utils.TrimFieldName(qi.Key)
			if checkRepositoryValue(r, sKey, qi.Value, qi.Aggregator) {
				if inclusive {
					return true
				}
			} else {
				if !inclusive {
					return false
				}
			}
		}
	}
	return false
}

func (rn *repositoryManager) filter(inclusive bool, q ...model.Query) []model.Repository {
	var out = make([]model.Repository, 0)
	for _, r := range rn.manager.GetRepositoryList() {
		if rn.checkFilter(r, inclusive, q...) {
			out = append(out, r)
		}
	}
	return out
}

func (rn *repositoryManager) DeleteRepositories(inclusive bool, q ...model.Query) model.DataResponse {
	var respObjs = make([]interface{}, 0)
	var list = rn.filter(inclusive, q...)
	var message = ""
	for _, r := range list {
		r.State = model.StateDeleted
		err := rn.manager.SaveRepository(r)
		if err != nil {
			if len(message) > 0 {
				message += ", "
			}
			message += fmt.Sprintf("repository: %s - Error: %v", r.Name, err)
		} else {
			respObjs = append(respObjs, r)
		}
	}
	var success = false
	if len(message) == 0 {
		success = true
		message = "OK"
	}
	return model.DataResponse{
		Success:         success,
		Message:         message,
		ResponseObjects: respObjs,
		Changes:         int64(len(list)),
	}
}

func (rn *repositoryManager) PurgeRepositories(inclusive bool, q ...model.Query) model.DataResponse {
	var respObjs = make([]interface{}, 0)
	var list = rn.filter(inclusive, q...)
	var message = ""
	for _, r := range list {
		err := rn.manager.DeleteRepositoryById(r.Id)
		if err != nil {
			if len(message) > 0 {
				message += ", "
			}
			message += fmt.Sprintf("repository: %s - Error: %v", r.Name, err)
		} else {
			respObjs = append(respObjs, r)
		}
	}
	var success = false
	if len(message) == 0 {
		success = true
		message = "OK"
	}
	return model.DataResponse{
		Success:         success,
		Message:         message,
		ResponseObjects: respObjs,
		Changes:         int64(len(list)),
	}
}

func (rn *repositoryManager) ClearRepository(id string) model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (rn *repositoryManager) ClearRepositoryByName(name string) model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (rn *repositoryManager) GetRepository(id string) *model.Repository {
	return nil
}

func (rn *repositoryManager) GetRepositoryByName(name string) *model.Repository {
	return nil
}

func (rn *repositoryManager) AccessRepository(r model.Repository) *model.DocumentsDataManager {
	dm := GetDocumentDataManager(rn.baseDataFolder, &r)
	return &dm
}

func (rn *repositoryManager) OverrideRepository(id string, r *model.Repository) model.DataResponse {
	if r == nil {
		return model.DataResponse{
			Success:         false,
			Message:         fmt.Sprintf("Error overriding repository id : %s with empty body", id),
			ResponseObjects: nil,
		}
	}
	rp, err := rn.manager.OverrideRepository(id, *r)
	var response = make([]interface{}, 0)
	if err != nil {
		return model.DataResponse{
			Success:         false,
			Message:         fmt.Sprintf("Error updating repository id: %s, error: %v", id, err),
			ResponseObjects: response,
		}
	} else {
		response = append(response, *rp)
		return model.DataResponse{
			Success:         true,
			Message:         "REPOSITORY Overridden",
			ResponseObjects: response,
		}
	}
}
