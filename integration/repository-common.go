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

func loadCharts(dataFolder string, logger log.Logger, repoName string) ([]model.Chart, error) {
	var emptyChartList = make([]model.Chart, 0)
	var chartFileList = model.ChartList{
		RepoName: repoName,
		Charts:   emptyChartList,
	}
	// Load Repository Charts File
	var file = fmt.Sprintf(repositoryChartsIndexTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Loading charts file %s for repository %s", file, repoName)
	}
	var folder = fmt.Sprintf(repositoryChartsFolderTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator)
	if !utils.ExistsFileOrFolder(folder) {
		_ = utils.CleanCreateFolder(folder)
		if logger != nil {
			logger.Warnf("Folder %s already exists, reusing charts ...", folder)
		}
	}
	err := utils.LoadStructureByType(file, &chartFileList, repositoryFormatExtension)
	if err != nil {
		return emptyChartList, err
	}
	var chartList = chartFileList.Charts
	if logger != nil {
		logger.Warnf("Number of loaded charts %v for repository %s", len(chartList), repoName)
	}
	return chartList, nil
}

func saveKubernetesFiles(dataFolder string, logger log.Logger, repoName string, kubernetesFiles []model.KubernetesFile) error {
	var kubernetesFileList = model.KubeFileList{
		RepoName: repoName,
		Files:    kubernetesFiles,
	}
	// Create Repository Charts File
	var file = fmt.Sprintf(repositoryKubernetesFilesIndexTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Saving Kubernetes Files file %s for repository %s", file, repoName)
		logger.Warnf("Number of saved Kubernetes Files %v for repository %s", len(kubernetesFiles), repoName)
	}
	var folder = fmt.Sprintf(repositoryKubernetesFilesFolderTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator)
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

func loadKubernetesFiles(dataFolder string, logger log.Logger, repoName string) ([]model.KubernetesFile, error) {
	var emptyKubernetesFileList = make([]model.KubernetesFile, 0)
	var kubernetesFileList = model.KubeFileList{
		RepoName: repoName,
		Files:    emptyKubernetesFileList,
	}
	// Create Repository Charts File
	var file = fmt.Sprintf(repositoryKubernetesFilesIndexTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Loading Kubernetes Files file %s for repository %s", file, repoName)
	}
	var folder = fmt.Sprintf(repositoryKubernetesFilesFolderTemplate, dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator)
	if !utils.ExistsFileOrFolder(folder) {
		if logger != nil {
			logger.Warnf("Folder %s already exists, reusing Kubernetes files ...", folder)
		}
	}
	err := utils.LoadStructureByType(file, &kubernetesFileList, repositoryFormatExtension)
	if err != nil {
		return emptyKubernetesFileList, err
	}
	var filesList = kubernetesFileList.Files
	if logger != nil {
		logger.Warnf("Number of loaded Kubernetes Files %v for repository %s", len(filesList), repoName)
	}
	return filesList, nil
}
