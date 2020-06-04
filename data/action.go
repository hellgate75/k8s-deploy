package data

import (
	"github.com/hellgate75/go-services/database"
	"github.com/hellgate75/k8s-deploy/data/device"
	"github.com/hellgate75/k8s-deploy/data/mongo"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
)

func GetDeviceRepositoryDataManager(baseFolder string, manager model.RepositoryStorageManager, logger log.Logger) model.RepositoryDataManager {
	return device.GetRepositoryDataManager(baseFolder, manager, logger)
}

func GetMongoRepositoryDataManager(conn database.Connection, baseFolder string, manager model.RepositoryStorageManager, logger log.Logger) model.RepositoryDataManager {
	return mongo.GetRepositoryDataManager(conn, baseFolder, manager, logger)
}

func GetDeviceDocumentsDataManager(baseFolder string, repo *model.Repository) model.DocumentsDataManager {
	return device.GetDocumentDataManager(baseFolder, repo)
}

func GetMongoDocumentsDataManager(conn database.Connection, repo *model.Repository) model.DocumentsDataManager {
	return mongo.GetDocumentDataManager(conn, repo)
}
