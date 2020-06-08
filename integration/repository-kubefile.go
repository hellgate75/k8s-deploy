package integration

import (
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
)

type kubefileRepositoryManager struct {
	repository model.Repository
	dataFolder string
	logger     log.Logger
}

const (
	repositoryKubefilesIndexTemplate                = "%s%crepositories%c%s%ckubefiles%cindex.%v"
	repositoryKubefilesFolderTemplate               = "%s%crepositories%c%s%ckubefiles"
	repositoryKubefileDetailsIndexTemplate          = "%s%crepositories%c%s%ckubefiles%c%s%cindex.%v"
	repositoryKubefileDetailsFolderTemplate         = "%s%crepositories%c%s%ckubefiles%c%s"
	repositoryKubefileVersionsDetailsFolderTemplate = "%s%crepositories%c%s%ckubefiles%c%s%c%s"
)

func (k *kubefileRepositoryManager) VerifyKubeFile(name string, version string) error {
	panic("implement me")
}

func (k *kubefileRepositoryManager) InstallKubeFile(name string, version string, file string) error {
	panic("implement me")
}

func (k *kubefileRepositoryManager) DeleteKubeFileVersion(name string, version string) error {
	panic("implement me")
}

func (k *kubefileRepositoryManager) DeleteEntireKubeFile(name string, version string) error {
	panic("implement me")
}

func (k *kubefileRepositoryManager) GetKubeFileVersionTemplate(name string, version string) (string, error) {
	panic("implement me")
}

func (k *kubefileRepositoryManager) UpdateExistingKubeFile(name string, version string, file string) error {
	panic("implement me")
}

func (k *kubefileRepositoryManager) GetKubeFileVersions(name string) ([]model.Version, error) {
	panic("implement me")
}

func (k *kubefileRepositoryManager) GetKubeFileProjectVersions(name string) ([]model.ProjectKubeFile, error) {
	panic("implement me")
}

func (k *kubefileRepositoryManager) DeployInstallKubeFile(name string, version string) (string, error) {
	panic("implement me")
}

func (k *kubefileRepositoryManager) DeployUpgradeKubeFile(name string, version string, force bool) (string, error) {
	panic("implement me")
}

func (k *kubefileRepositoryManager) GetInstalledKubeFileVersion(name string) (model.Version, error) {
	panic("implement me")
}

func (k *kubefileRepositoryManager) GetInstalledKubeFileVersionDetails(name string, version string) (model.Version, error) {
	panic("implement me")
}

func (k *kubefileRepositoryManager) UndeployInstalledKubeFile(name string) (model.Version, error) {
	panic("implement me")
}

func NewRepositoryKubeFilesManager(repository model.Repository, dataFolder string, logger log.Logger) model.RepositoryKubeFilesManager {
	return &kubefileRepositoryManager{
		repository: repository,
		dataFolder: dataFolder,
		logger:     logger,
	}
}
