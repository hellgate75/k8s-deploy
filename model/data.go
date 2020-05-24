package model

// Represents the data interaction response structure
type DataResponse struct {
	// Success state flag
	Success			bool
	// Error message
	Message			string
	// Number of changes in the data operation
	Changes			int64
	// Created, Modified or Deleted objects
	ResponseObjects	[]interface{}
}

// Represents the k8srepo documents data storage manager
type DocumentsDataManager interface {
	// Add new chart data
	AddChart(c Chart) DataResponse
	// Add new Kubefile data
	AddKubeFile(f KubeFile) DataResponse
	// Add new chart version data
	AddChartVersion(c Chart, v Version) DataResponse
	// Add new Kubefile version data
	AddKubeFileVersion(f KubeFile, v Version) DataResponse
	// Remove one or more Charts
	RemoveCharts(q ...Query) DataResponse
	// Remove one or more Kube-files
	RemoveKubeFiles(q ...Query) DataResponse
	// Remove one or more Chart versions
	RemoveChartVersions(c Chart, q ...Query) DataResponse
	// Remove one or more Kube-file versions
	RemoveKubeFileVersions(f KubeFile, q ...Query) DataResponse
	// Purge permanently one or more Charts
	PurgeCharts(q ...Query) DataResponse
	// Purge permanently one or more Kube-files
	PurgeKubeFiles(q ...Query) DataResponse
	// Purge permanently one or more Chart versions
	PurgeChartVersions(c Chart, q ...Query) DataResponse
	// Purge permanently one or more Kube-file versions
	PurgeKubeFileVersions(f KubeFile, q ...Query) DataResponse
	// Update an existing Charts
	UpdateCharts(c Chart, q ...Query) DataResponse
	// Update one or more Kube-files
	UpdateKubeFiles(f KubeFile, v Version, q ...Query) DataResponse
	// Update an existing Chart versions
	UpdateChartVersions(c Chart, v Version, q ...Query) DataResponse
	// Update one or more Kube-file versions
	UpdateKubeFileVersions(f KubeFile, v Version, q ...Query) DataResponse
	// Query over k8srepo Charts
	QueryCharts(q ...Query) DataResponse
	// Query over k8srepo Kube-files
	QueryKubeFiles(q ...Query) DataResponse
	// Query over k8srepo Chart versions
	QueryChartVersions(c Chart, q ...Query) DataResponse
	// Query over k8srepo Kube-file versions
	QueryKubeFileVersions(f KubeFile, q ...Query) DataResponse
	// List all k8srepo Charts
	ListCharts() DataResponse
	// List all k8srepo Kube-files
	ListKubeFiles() DataResponse
	// List k8srepo Chart versions
	ListChartVersions(q ...Query) DataResponse
	// List all k8srepo Kube-file versions
	ListKubeFileVersions(q ...Query) DataResponse
}

// Represents the repositories data storage manager
type RepositoryDataManager interface {
	//List all repositories
	ListRepositories() DataResponse
	//Add new k8srepo
	AddRepository(r Repository) DataResponse
	//Delete one or more repositories
	DeleteRepositories(q ...Query) DataResponse
	//Purge data for one or more repositories
	PurgeRepositories(q ...Query) DataResponse
	//Clear a repositories data entities selecting by id
	ClearRepository(id string) DataResponse
	//Clear a repositories data entities selecting by name
	ClearRepositoryByName(name string) DataResponse
	//Collect a k8srepo data entities selecting by id
	GetRepository(id string) *Repository
	//Collect a k8srepo data entities selecting by name
	GetRepositoryByName(name string) *Repository
	//Access Repository data via model.DocumentsDataManager
	AccessRepository(r Repository) *DocumentsDataManager
	//Override a k8srepo selecting via id
	OverrideRepository(id string, r Repository) DataResponse
}

// Represents the projects data storage manager
type ProjectDataManager interface {
	// List all projects
	ListProjects() DataResponse
	// Add a new project
	AddProject(p Project) DataResponse
	// Add a new version to a project
	AddProjectVersion(p Project, version ProjectVersion) DataResponse
	// Delete one ot more projects
	DeleteProjects(q ...Query) DataResponse
	// Delete one ot more project verions
	DeleteProjectVersions(p Project, q ...Query) DataResponse
	// Purge one ot more projects definitely
	PurgeProjects(q ...Query) DataResponse
	// Purge one ot more project versions definitely
	PurgeProjectVersions(p Project, q ...Query) DataResponse
	// Clear project information, selecting by id
	ClearProject(id string) DataResponse
	// Clear project information, selecting by name
	ClearProjectByName(name string) DataResponse
	// Collect project information, selecting by id
	GetProject(id string) *Project
	// Collect project information, selecting by name
	GetProjectByName(name string) *Project
	// Collect project versions information
	GetProjectVersions(p Project, q ...Query) *Project
	// Override project information
	OverrideProject(id string, p Project) DataResponse
	// Override project versions information
	OverrideProjectVersions(id string, v Version, q ...Query) DataResponse
	// Query over Projects
	QueryProjects(q ...Query) DataResponse
	// Query over Project versions
	QueryProjectVersions(c Chart, q ...Query) DataResponse
}


// Represents the deploy jobs data storage manager
type JobsDataManager interface {
	//List all Deploy jobs
	ListJobs() DataResponse
	//Add new job to the Deploy
	AddJob(j Job) DataResponse
	//Delete one or jobs in a Deploy
	DeleteJobs(q ...Query) DataResponse
	//Purge definitely one or jobs in a Deploy
	PurgeJobs(q ...Query) DataResponse
	//Collect job data by job id
	GetJob(id string) *Job
	//Collect job data by job name
	GetJobByName(name string) *Job
	//Override existing job data in the Deploy
	OverrideJob(id string, j Job) DataResponse
}

// Represents the deploys data storage manager
type DeployDataManager interface {
	// List all Deploys
	ListDeploys() DataResponse
	// Add new Deploy
	AddDeploy(d Deploy) DataResponse
	// Delete existing Deploy
	DeleteDeploys(q ...Query) DataResponse
	// Purge one or more Deploys
	PurgeDeploys(q ...Query) DataResponse
	// Clear Deploys data selecting by id
	ClearDeploy(id string) DataResponse
	// Clear Deploys data selecting by name
	ClearDeployByName(name string) DataResponse
	// Get Deploy selecting by id
	GetDeploy(id string) *Deploy
	// Get Deploy selecting by name
	GetDeployByName(name string) *Deploy
	// Access to deploy jobs data using model.JobsDataManager
	AccessDeploy(d Deploy) *JobsDataManager
	// Override existing Deploy selecting by Id
	OverrideDeploy(id string, d Deploy) DataResponse
}

// Represents the global data storage manager
type DataManager struct {
	// Manage Repositories Data
	Repos 		RepositoryDataManager
	// Manage Projects Data
	Projects	ProjectDataManager
	// Manage Deploys Data
	Deploys		DeployDataManager
}
