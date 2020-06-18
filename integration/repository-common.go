package integration

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
	"os"
)

func getChartsListFolder(baseFolder string, repoName string) string {
	return fmt.Sprintf(repositoryChartsFolderTemplate, baseFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator)
}

func getChartsListIndex(baseFolder string, repoName string, extension utils.FormatType) string {
	//repositoryChartsIndexTemplate                = "%s%crepositories%c%s%ccharts%cindex.%v"
	return fmt.Sprintf(repositoryChartsIndexTemplate, baseFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, extension)
}

func getChartDetailsFolder(baseFolder string, repoName string, chartName string) string {
	//repositoryChartDetailsFolderTemplate         = "%s%crepositories%c%s%ccharts%c%s"
	return fmt.Sprintf(repositoryChartDetailsFolderTemplate, baseFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, chartName)
}

func getChartDetailsIndex(baseFolder string, repoName string, chartName string, extension utils.FormatType) string {
	//repositoryChartDetailsIndexTemplate          = "%s%crepositories%c%s%ccharts%c%s%cindex.%v"
	return fmt.Sprintf(repositoryChartDetailsIndexTemplate, baseFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, chartName, os.PathSeparator, extension)
}

func getChartVersionFolder(baseFolder string, repoName string, chartName string, version string) string {
	//repositoryChartVersionsDetailsFolderTemplate = "%s%crepositories%c%s%ccharts%c%s%c%s"
	return fmt.Sprintf(repositoryChartDetailsFolderTemplate, baseFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, os.PathSeparator, chartName, os.PathSeparator, version)
}

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

func saveCharts(dataFolder string, logger log.Logger, repoName string, charts []model.ChartInfo) error {
	var chartFileList = model.ChartList{
		RepoName: repoName,
		Charts:   charts,
	}
	// Create Repository Charts File
	var file = getChartsListIndex(dataFolder, repoName, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Saving charts file %s for repository %s", file, repoName)
		logger.Warnf("Number of saved charts %v for repository %s", len(charts), repoName)
	}
	var folder = getChartsListFolder(dataFolder, repoName)
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

func loadCharts(dataFolder string, logger log.Logger, repoName string) ([]model.ChartInfo, error) {
	var emptyChartList = make([]model.ChartInfo, 0)
	var chartFileList = model.ChartList{
		RepoName: repoName,
		Charts:   emptyChartList,
	}
	// Load Repository Charts File
	var file = getChartsListIndex(dataFolder, repoName, repositoryFormatExtension)
	if logger != nil {
		logger.Warnf("Loading charts file %s for repository %s", file, repoName)
	}
	var folder = getChartsListFolder(dataFolder, repoName)
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

func saveKubernetesFiles(dataFolder string, logger log.Logger, repoName string, kubernetesFiles []model.KubernetesFileInfo) error {
	var kubernetesFileList = model.KubernetesFileList{
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

func loadKubernetesFiles(dataFolder string, logger log.Logger, repoName string) ([]model.KubernetesFileInfo, error) {
	var emptyKubernetesFileList = make([]model.KubernetesFileInfo, 0)
	var kubernetesFileList = model.KubernetesFileList{
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
