package mongo

import (
	"fmt"
	"github.com/hellgate75/go-services/database"
	"github.com/hellgate75/k8s-deploy/model"
	"os"
)

func getRepositoryFolder(basePath string, repoName string) string {
	return fmt.Sprintf("%s%crepos%c%s", basePath, os.PathSeparator, os.PathSeparator, repoName)
}

type repositoryManager struct {
	conn	*database.Connection
}

func (rn *repositoryManager) ListRepositories()  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (rn *repositoryManager) AddRepository(r model.Repository)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (rn *repositoryManager) DeleteRepositories(q ... model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (rn *repositoryManager) PurgeRepositories(q ... model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (rn *repositoryManager) ClearRepository(id string)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (rn *repositoryManager) ClearRepositoryByName(name string)  model.DataResponse {
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
	dm := getDocumentDataManager(rn.baseDataFolder, &r)
	return &dm
}


func (rn *repositoryManager) OverrideRepository(id string, r model.Repository)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func getRepositoryDataManager(conn *database.Connection) model.RepositoryDataManager {
	return &repositoryManager {
		conn: conn,
	}
}

