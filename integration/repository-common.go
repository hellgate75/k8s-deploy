package integration

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
	"os"
)

func saveRepository(dataFolder string, logger log.Logger, repoName string, repo model.Repository) error {
	// Create Repository files
	var file = fmt.Sprintf(repositoryDetailsIndexTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Saving Repository file %s for repository %s", file, repoName)
	}
	var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName)
	if !utils.ExistsFileOrFolder(folder) {
		err := utils.CleanCreateFolder(folder)
		if err != nil {
			return err
		}
	} else {
		if logger != nil {
			logger.Warnf("Folder %s already exists, trying existance of charts and kubefiles ...", folder)
		}
	}
	return utils.SaveStructureByType(file, &repo, repositoryFormatExtension)
}

func saveCharts(dataFolder string, logger log.Logger, repoName string, charts []model.Chart) error {
	var chartFileList = model.ChartList{
		RepoName: repoName,
		Charts:   charts,
	}
	// Create Repository Charts File
	var file = fmt.Sprintf(repositoryChartsIndexTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Saving charts file %s for repository %s", file, repoName)
		logger.Warnf("Number of saved charts %v for repository %s", len(charts), repoName)
	}
	var folder = fmt.Sprintf(repositoryChartsFolderTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator)
	if !utils.ExistsFileOrFolder(folder) {
		err := utils.CleanCreateFolder(folder)
		if err != nil {
			return err
		}
	} else {
		if logger != nil {
			logger.Warnf("Folder %s already exists, reusing charts ...", folder)
		}
	}
	return utils.SaveStructureByType(file, &chartFileList, repositoryFormatExtension)
}

func saveKubernetesFiles(dataFolder string, logger log.Logger, repoName string, kubernetesFiles []model.KubernetesFile) error {
	var kubernetesFileList = model.KubeFileList{
		RepoName: repoName,
		Files:    make([]model.KubernetesFile, 0),
	}
	// Create Repository Charts File
	var file = fmt.Sprintf(repositoryKubefilesIndexTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Saving Kubernetes Files file %s for repository %s", file, repoName)
		logger.Warnf("Number of saved Kubernetes Files %v for repository %s", len(kubernetesFiles), repoName)
	}
	var folder = fmt.Sprintf(repositoryKubefilesFolderTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator)
	if !utils.ExistsFileOrFolder(folder) {
		err := utils.CleanCreateFolder(folder)
		if err != nil {
			return err
		}
	} else {
		if logger != nil {
			logger.Warnf("Folder %s already exists, reusing kubenretes files ...", folder)
		}
	}
	return utils.SaveStructureByType(file, &kubernetesFileList, repositoryFormatExtension)
}
