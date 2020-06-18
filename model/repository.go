package model

//Describes the Repository Charts Manager interface
type RepositoryStorageManager interface {
	// Gets the list of existing repositories
	GetRepositoryList() []Repository
	// save an existing repository
	SaveRepository(repository Repository) error
	//Crete a new named repository, if the name is not in use yet
	CreateRepository(name string) (*Repository, error)
	//Update an existing repository or create a new named one, if the id is reflects to an existing reporisotry
	UpdateRepository(id string, r Repository) (*Repository, error)
	//Override an existing repository or create a new named one, if the id is reflects to an existing reporisotry
	OverrideRepository(id string, r Repository) (*Repository, error)
	// Delete a repository using the name
	DeleteRepositoryByName(name string) error
	// Delete a repository using the identifier
	DeleteRepositoryById(name string) error
	// Rename a given repository
	RenameRepository(oldName string, newName string) error
	// List all repository charts
	ListRepositoryCharts(id string) ([]ChartInfo, error)
	// List all repository Kubernetes yaml files
	ListRepositoryKubernetesFiles(id string) ([]KubernetesFileInfo, error)
	// Backup an existing repository
	BackupRepository(id string, archiveFile string, useZipFormat bool) error
	// Restore an existing repository from zip/tar archive
	RestoreRepository(archiveFile string, useZipFormat bool, forceCreate bool) error
	// Gets Charts Manager for given repository
	GetRepositoryChartsManager(id string) (RepositoryChartManager, error)
	// Gets Kubernetes yaml files Manager for given repository
	GetRepositoryKubernetesFilesManager(id string) (RepositoryKubernetesFilesManager, error)
	// Gets Charts Manager for given repository
	GetRepositoryChartsManagerByName(id string) (RepositoryChartManager, error)
	// Gets Kubernetes yaml files Manager for given repository
	GetRepositoryKubernetesFilesManagerByName(id string) (RepositoryKubernetesFilesManager, error)
	// Initializes and loads all repositories
	Initialize() (RepositoryStorageManager, error)
	// Reloads all repositories information from storage
	Refresh() error
	// Saves current repository stage
	SavePoint() error
	// Get a single repository by name
	GetRepository(name string) (*Repository, error)
	// Get a single repository by id
	GetRepositoryById(id string) (*Repository, error)
}

//Describes the Repository Charts Manager interface
type RepositoryChartManager interface {
	//Verify presence and correctness of a chart
	VerifyChart(name string, version string) error
	//Install chart version via archive
	InstallChart(name string, version string, archive string, zipArchive bool) error
	// Delete a version of the Chart
	DeleteChartVersion(name string, version string) error
	// Delete an entire Chart, including all versions
	DeleteEntireChart(name string, version string) error
	// Get a yaml template build from a Chart version
	GetChartVersionTemplate(name string, version string, values ValueSet) (string, error)
	// Update and existing Chart, or if force create a new chart
	UpdateExistingChart(name string, version string, archive string, zipArchive bool, forceCreate bool) error
	// Collects versions of a Chart
	GetChartVersions(name string) ([]Version, error)
	// Collects project versions of a Chart, ready for job scheduling
	GetChartProjectVersions(name string) ([]ProjectChart, error)
	// Execute deploy of a chart and collects the output
	DeployInstallChart(name string, version string, values ValueSet) (string, error)
	// Execute upgrade of a chart and collects the output
	DeployUpgradeChart(name string, version string, values ValueSet, force bool) (string, error)
	// Verify and return chart version, or an error in case chart is not installed
	GetInstalledChartVersion(name string) (Version, error)
	// Get isntalled Chart version details
	GetInstalledChartVersionDetails(name string, version string) (Version, error)
	// Un-deploy installed chart, and collects latest installed version
	UnDeployInstalledChart(name string) (Version, error)
}

//Describes the Repository Kubernetes yaml files Manager interface
type RepositoryKubernetesFilesManager interface {
	//Verify presence and correctness of a Kubernetes yaml file
	VerifyKubernetesFile(name string, version string) error
	//Install Kubernetes yaml file version via file
	InstallKubernetesFile(name string, version string, file string) error
	// Delete a version of the Kubernetes yaml file
	DeleteKubernetesFileVersion(name string, version string) error
	// Delete an entire Kubernetes yaml file, including all versions
	DeleteEntireKubernetesFile(name string, version string) error
	// Get a yaml template build from a Kubernetes yaml file version
	GetKubernetesFileVersionTemplate(name string, version string) (string, error)
	// Update and existing Kubernetes yaml file, or createa new Kubernetes yaml file deployment
	UpdateExistingKubernetesFile(name string, version string, file string) error
	// Collects versions of a Kubernetes yaml file
	GetKubernetesFileVersions(name string) ([]Version, error)
	// Collects project versions of a Kubernetes yaml file, ready for job scheduling
	GetKubernetesFileProjectVersions(name string) ([]ProjectKubeFile, error)
	// Execute deploy of a Kubernetes yaml file and collects the output
	DeployInstallKubernetesFile(name string, version string) (string, error)
	// Execute upgrade of a Kubernetes yaml file and collects the output
	DeployUpgradeKubernetesFile(name string, version string, force bool) (string, error)
	// Verify and return Kubernetes yaml file version, or an error in case Kubernetes yaml file is not installed
	GetInstalledKubernetesFileVersion(name string) (Version, error)
	// Get installed Kubernetes yaml file version details
	GetInstalledKubernetesFileVersionDetails(name string, version string) (Version, error)
	// Un-deploy installed Kubernetes yaml file, and collects latest installed version
	UnDeployInstalledKubernetesFile(name string) (Version, error)
}
