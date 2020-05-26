package integration

import (
	"errors"
	"fmt"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
	"os"
	"sync"
)

type repositoryStorageManager struct {
	dataFolder string
}

func (s *repositoryStorageManager) GetRepositoriesList() []model.Repository {
	panic("implement me")
}

func (s *repositoryStorageManager) CreateRepository(name string) (model.Repository, error) {
	panic("implement me")
}

func (s *repositoryStorageManager) DeleteRepositoryByName(name string) error {
	panic("implement me")
}

func (s *repositoryStorageManager) DeleteRepositoryById(name string) error {
	panic("implement me")
}

func (s *repositoryStorageManager) RenameRepository(oldName string, newName string) error {
	panic("implement me")
}

func (s *repositoryStorageManager) ListRepositoryCharts(id string) ([]model.Chart, error) {
	panic("implement me")
}

func (s *repositoryStorageManager) ListRepositoryKubeFiles(id string) ([]model.KubeFile, error) {
	panic("implement me")
}

func (s *repositoryStorageManager) BackupRepository(id string, archiveFile string, useZipFormat bool) error {
	panic("implement me")
}

func (s *repositoryStorageManager) RestoreRepository(id string, archiveFile string, useZipFormat bool, forceCreate bool) error {
	panic("implement me")
}

func (s *repositoryStorageManager) GetRepositoryMutex(id string) (sync.RWMutex, error) {
	panic("implement me")
}

func (s *repositoryStorageManager) GetRepositoryChartsManager(id string) (model.RepositoryChartManager, error) {
	panic("implement me")
}

func (s *repositoryStorageManager) GetRepositoryKubeFilesManager(id string) (model.RepositoryKubeFilesManager, error) {
	panic("implement me")
}

const (
	repositoryIndeTemplate    = "%s%crepositories.%v"
	repositoryFormatExtension = utils.YAML_FORMAT
)

func (s *repositoryStorageManager) Initialize() (model.RepositoryStorageManager, error) {
	if _, err := os.Stat(s.dataFolder); err != nil {
		err = os.MkdirAll(s.dataFolder, 666)
		if err != nil {
			return nil, err
		}
	}
	var file = fmt.Sprintf(repositoryIndeTemplate, s.dataFolder, os.PathSeparator, repositoryFormatExtension)
	if utils.ExistsFileOrFolder(file) {
		return s, s.Refresh()
	}
	return s, nil
}

func (s *repositoryStorageManager) SavePoint() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var file = fmt.Sprintf(repositoryIndeTemplate, s.dataFolder, os.PathSeparator, repositoryFormatExtension)
	if utils.ExistsFileOrFolder(file) {
		_ = utils.DeleteFileOrFolder(file)
	}
	err = utils.SaveStructureByType(file, s, repositoryFormatExtension)
	return err
}

func (s *repositoryStorageManager) Refresh() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var file = fmt.Sprintf(repositoryIndeTemplate, s.dataFolder, os.PathSeparator, repositoryFormatExtension)
	err = utils.LoadStructureByType(file, s, repositoryFormatExtension)
	return err
}

var repositoryStorageManagerSingleton model.RepositoryStorageManager

func GetRepositoryStorageManagerSingleton(dataFolder string) (model.RepositoryStorageManager, error) {
	var err error
	if repositoryStorageManagerSingleton == nil {
		repositoryStorageManagerSingleton, err = (&repositoryStorageManager{
			dataFolder: dataFolder,
		}).Initialize()
		if err != nil {
			repositoryStorageManagerSingleton = nil
			return nil, err
		}
	}
	return repositoryStorageManagerSingleton, err
}

func NewRepositoryStorageManager(dataFolder string) (model.RepositoryStorageManager, error) {
	return (&repositoryStorageManager{
		dataFolder: dataFolder,
	}).Initialize()
}
