package integration

import (
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
)

type kubernetesFilesRepositoryManager struct {
	repository model.Repository
	dataFolder string
	logger     log.Logger
	files      []model.KubernetesFileInfo
}

const (
	repositoryKubernetesFilesIndexTemplate                = "%s%crepositories%c%s%ckubefiles%cindex.%v"
	repositoryKubernetesFilesFolderTemplate               = "%s%crepositories%c%s%ckubefiles"
	repositoryKubernetesFileDetailsIndexTemplate          = "%s%crepositories%c%s%ckubefiles%c%s%cindex.%v"
	repositoryKubernetesFileDetailsFolderTemplate         = "%s%crepositories%c%s%ckubefiles%c%s"
	repositoryKubernetesFileVersionsDetailsFolderTemplate = "%s%crepositories%c%s%ckubefiles%c%s%c%s"
)

func (k *kubernetesFilesRepositoryManager) VerifyKubernetesFile(name string, version string) error {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) InstallKubernetesFile(name string, version string, file string) error {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) DeleteKubernetesFileVersion(name string, version string) error {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) DeleteEntireKubernetesFile(name string, version string) error {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) GetKubernetesFileVersionTemplate(name string, version string) (string, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) UpdateExistingKubernetesFile(name string, version string, file string) error {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) GetKubernetesFileVersions(name string) ([]model.Version, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) GetKubernetesFileProjectVersions(name string) ([]model.ProjectKubeFile, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) DeployInstallKubernetesFile(name string, version string) (string, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) DeployUpgradeKubernetesFile(name string, version string, force bool) (string, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) GetInstalledKubernetesFileVersion(name string) (model.Version, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) GetInstalledKubernetesFileVersionDetails(name string, version string) (model.Version, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) UnDeployInstalledKubernetesFile(name string) (model.Version, error) {
	panic("implement me")
}

func (k *kubernetesFilesRepositoryManager) init() (model.RepositoryKubernetesFilesManager, error) {
	files, err := loadKubernetesFiles(k.dataFolder, k.logger, k.repository.Name)
	if err != nil {
		return k, err
	}
	k.files = files
	return k, nil
}

func NewRepositoryKubernetesFilesManager(repository model.Repository, dataFolder string, logger log.Logger) (model.RepositoryKubernetesFilesManager, error) {
	return (&kubernetesFilesRepositoryManager{
		repository: repository,
		dataFolder: dataFolder,
		logger:     logger,
		files:      make([]model.KubernetesFileInfo, 0),
	}).init()
}
