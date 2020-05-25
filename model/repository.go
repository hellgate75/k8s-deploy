package model

//Represents the Repository Charts manager
type RepositoryChartManager interface {
	//Verify presence and correctness of a chart
	VerifyChart(name string, version string) error
	//Install chart version via archive
	InstallChart(name string, version string, archive string, zipArchive bool) error
	// Delete a version of the Chart
	DeleteChartVersion(name string, version string) error
	// Delete an entire Chart, including all versions
	DeleteEntireChart(name string, version string) error
	// Delete a version of the Chart
	GetChartVersionTemplate(name string, version string) (string, error)
	// Update and existing Chart, or if force create a new chart
	UpdateExistingChart(name string, version string, archive string, zipArchive bool, forceCreate bool) error
	// Collects versions of a Chart
	GetChartVersions(name string) ([]Version, error)
	// Collects project versions of a Chart, ready for job scheduling
	GetChartProjectVersions(name string) ([]ProjectChart, error)
	// Execute deploy of a chart and collects the output
	ExecuteChartDeployCommand(name string, version string) (string, error)
	// Execute upgrade of a chart and collects the output
	ExecuteChartUpgradeCommand(name string, version string, force bool) (string, error)
	// Verify and return chart version, or an error in case chart is not installed
	GetInstalledChartVersion(name string) (Version, error)
	// Get isntalled Chart version details
	GetInstalledChartVersionDetails(name string, version string) (Version, error)
	// Un-deploy installed chart, and collects latest installed version
	UndeployInstalledChart(name string) (Version, error)

}
