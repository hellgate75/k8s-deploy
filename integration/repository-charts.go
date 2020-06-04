package integration

import (
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
)

type chartsRepositoryManager struct {
	repository model.Repository
	dataFolder string
	logger     log.Logger
}

const (
	repositoryChartsIndexTemplate  = "%s%crepositories%c%s%ccharts%cindex.%v"
	repositoryChartsFolderTemplate = "%s%crepositories%c%s%ccharts"
)

func (c *chartsRepositoryManager) VerifyChart(name string, version string) error {
	panic("implement me")
}

func (c *chartsRepositoryManager) InstallChart(name string, version string, archive string, zipArchive bool) error {
	panic("implement me")
}

func (c *chartsRepositoryManager) DeleteChartVersion(name string, version string) error {
	panic("implement me")
}

func (c *chartsRepositoryManager) DeleteEntireChart(name string, version string) error {
	panic("implement me")
}

func (c *chartsRepositoryManager) GetChartVersionTemplate(name string, version string, values model.ValueSet) (string, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) UpdateExistingChart(name string, version string, archive string, zipArchive bool, forceCreate bool) error {
	panic("implement me")
}

func (c *chartsRepositoryManager) GetChartVersions(name string) ([]model.Version, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) GetChartProjectVersions(name string) ([]model.ProjectChart, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) DeployInstallChart(name string, version string, values model.ValueSet) (string, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) DeployUpgradeChart(name string, version string, values model.ValueSet, force bool) (string, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) GetInstalledChartVersion(name string) (model.Version, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) GetInstalledChartVersionDetails(name string, version string) (model.Version, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) UndeployInstalledChart(name string) (model.Version, error) {
	panic("implement me")
}

func NewRepositoryChartManager(repository model.Repository, dataFolder string, logger log.Logger) model.RepositoryChartManager {
	return &chartsRepositoryManager{
		repository: repository,
		dataFolder: dataFolder,
		logger:     logger,
	}
}
