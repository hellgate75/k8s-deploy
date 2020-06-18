package integration

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
)

type chartsRepositoryManager struct {
	repository model.Repository
	dataFolder string
	logger     log.Logger
	charts     []model.ChartInfo
}

const (
	repositoryChartDetailsIndexTemplate          = "%s%crepositories%c%s%ccharts%c%s%cindex.%v"
	repositoryChartDetailsFolderTemplate         = "%s%crepositories%c%s%ccharts%c%s"
	repositoryChartVersionsDetailsFolderTemplate = "%s%crepositories%c%s%ccharts%c%s%c%s"
	repositoryChartsIndexTemplate                = "%s%crepositories%c%s%ccharts%cindex.%v"
	repositoryChartsFolderTemplate               = "%s%crepositories%c%s%ccharts"
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

func getChartIndexFile(baseFolder string, chartName string) string {
	return fmt.Sprintf(repositoryChartsFolderTemplate, baseFolder, chartName)
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

func (c *chartsRepositoryManager) UnDeployInstalledChart(name string) (model.Version, error) {
	panic("implement me")
}

func (c *chartsRepositoryManager) init() (model.RepositoryChartManager, error) {
	charts, err := loadCharts(c.dataFolder, c.logger, c.repository.Name)
	if err != nil {
		return c, err
	}
	c.charts = charts
	return c, nil
}

func NewRepositoryChartManager(repository model.Repository, dataFolder string, logger log.Logger) (model.RepositoryChartManager, error) {
	return (&chartsRepositoryManager{
		repository: repository,
		dataFolder: dataFolder,
		logger:     logger,
		charts:     make([]model.ChartInfo, 0),
	}).init()
}
