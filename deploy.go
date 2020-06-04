package k8s_deploy

import (
	"github.com/hellgate75/go-services/database"
	"github.com/hellgate75/k8s-deploy/data"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
	"net/http"
)

func GetDeviceRepositoryDataManager(baseFolder string, manager model.RepositoryStorageManager, logger log.Logger) model.RepositoryDataManager {
	return data.GetDeviceRepositoryDataManager(baseFolder, manager, logger)
}

func GetMongoRepositoryDataManager(conn database.Connection, baseFolder string, manager model.RepositoryStorageManager, logger log.Logger) model.RepositoryDataManager {
	return data.GetMongoRepositoryDataManager(conn, baseFolder, manager, logger)
}

func GetDeviceDocumentsDataManager(baseFolder string, repo *model.Repository) model.DocumentsDataManager {
	return data.GetDeviceDocumentsDataManager(baseFolder, repo)
}

func GetMongoDocumentsDataManager(conn database.Connection, repo *model.Repository) model.DocumentsDataManager {
	return data.GetMongoDocumentsDataManager(conn, repo)
}

func NewLogger(appName string, verbosity log.LogLevel) log.Logger {
	return log.NewLogger(appName, verbosity)
}

func NewFileLogger(appName string, logRotator log.LogRotator, verbosity log.LogLevel) (log.Logger, error) {
	return log.NewFileLogger(appName, logRotator, verbosity)
}

func SaveConfig(path string, name string, config interface{}) error {
	return model.SaveConfig(path, name, config)
}

func LoadConfig(path string, name string, config interface{}) error {
	return model.LoadConfig(path, name, config)
}

func RestParseRequest(w http.ResponseWriter, r *http.Request, res interface{}) error {
	return utils.RestParseResponse(w, r, res)
}

func RestParseResponse(w http.ResponseWriter, r *http.Request, req interface{}) error {
	return utils.RestParseResponse(w, r, req)
}
