package device

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	model2 "github.com/hellgate75/k8s-deploy/utils/model"
)

func GetRepositoryDataManager(baseFolder string, manager model.RepositoryStorageManager, logger log.Logger) model.RepositoryDataManager {
	return &repositoryManager{
		baseDataFolder: baseFolder,
		manager:        manager,
		logger:         logger,
	}
}

func checkRepositoryValue(r model.Repository, key string, value string, cond model.Aggregator) bool {
	switch key {
	case "name":
		return model2.CompareValues(value, r.Name, model2.DataTypeString, cond)
	case "id":
		return model2.CompareValues(value, r.Id, model2.DataTypeString, cond)
	case "state":
		return model2.CompareValues(value, fmt.Sprintf("%v", r.State), model2.DataTypeNumber, cond)
	case "charts":
		return model2.CompareValues(value, fmt.Sprintf("%v", len(r.GetCharts())), model2.DataTypeNumber, cond)
	case "kubernetesFiles":
		return model2.CompareValues(value, fmt.Sprintf("%v", len(r.GetKubernetesFiles())), model2.DataTypeNumber, cond)
	}
	return false
}
