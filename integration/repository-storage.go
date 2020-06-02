package integration

import (
	"errors"
	"fmt"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	repositoryIndeTemplate          = "%s%crepositories.%v"
	repositoryFormatExtension       = utils.YAML_FORMAT
	defaultRepositoryName           = "__default"
	repositoryDetailsIndexTemplate  = "%s%crepositories%c%s%cindex.%v"
	repositoryDetailsFolderTemplate = "%s%crepositories%c%s"
)

type RepositoryRef struct {
	Id   string
	Name string
}

type Repositories struct {
	Repositories []RepositoryRef `yaml:"repositories" json:"repositories" xml:"repository"`
	Created      time.Time       `yaml:"created" json:"created" xml:"created"`
	Updated      time.Time       `yaml:"updated" json:"updated" xml:"updated"`
}

type repositoryStorageManager struct {
	sync.RWMutex
	dataFolder   string
	repositories *Repositories
}

func (s *repositoryStorageManager) GetRepositoryList() []model.Repository {
	var out = make([]model.Repository, 0)
	for _, ref := range s.repositories.Repositories {
		repo, err := s.GetRepositoryById(ref.Id)
		if err == nil {
			out = append(out, *repo)
		} else {
			//TODO: Manage logging of error
		}
	}
	return out
}

func (s *repositoryStorageManager) containsRepositoryName(name string) bool {
	var repoName = utils.ConvertName(name)
	for _, repo := range s.repositories.Repositories {
		if repo.Name == repoName {
			return true
		}
	}
	return false
}

func (s *repositoryStorageManager) containsRepositoryId(id string) bool {
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			return true
		}
	}
	return false
}

func (s *repositoryStorageManager) getDefefaultRepositoryId() string {
	for _, repo := range s.repositories.Repositories {
		if repo.Name == defaultRepositoryName {
			return repo.Id
		}
	}
	return ""
}

func (s *repositoryStorageManager) GetRepository(name string) (*model.Repository, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.RUnlock()
	}()
	s.RLock()
	var repoName = utils.ConvertName(name)
	for _, repo := range s.repositories.Repositories {
		if repo.Name == repoName {
			var file = fmt.Sprintf(repositoryDetailsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, repositoryFormatExtension)
			var repository = model.Repository{}
			err := utils.LoadStructureByType(file, &repository, repositoryFormatExtension)
			if err != nil {
				return nil, err
			}
			return &repository, nil
		}
	}
	err = errors.New(fmt.Sprintf("Repository named: %s not found!!", name))
	return nil, err
}

func (s *repositoryStorageManager) GetRepositoryById(id string) (*model.Repository, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.RUnlock()
	}()
	s.RLock()
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			var file = fmt.Sprintf(repositoryDetailsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, repositoryFormatExtension)
			var repository = model.Repository{}
			err := utils.LoadStructureByType(file, &repository, repositoryFormatExtension)
			if err != nil {
				return nil, err
			}
			return &repository, nil
		}
	}
	err = errors.New(fmt.Sprintf("Repository with id: %s not found!!", id))
	return nil, err
}

func (s *repositoryStorageManager) CreateRepository(name string) (*model.Repository, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.Unlock()
	}()
	s.Lock()
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("Repository name cannot be empty or without significant digits or letters")
	}
	if s.containsRepositoryName(name) {
		return nil, errors.New(fmt.Sprintf("Repository name %s already present", name))
	}
	var repoName = utils.ConvertName(name)
	var newRepo = model.Repository{
		Name:      repoName,
		Id:        utils.NewUID(),
		State:     model.StateCreated,
		Charts:    make([]model.Chart, 0),
		KubeFiles: make([]model.KubeFile, 0),
	}
	var file = fmt.Sprintf(repositoryDetailsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repoName, os.PathSeparator, repositoryFormatExtension)
	var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repoName)
	if !utils.ExistsFileOrFolder(folder) {
		err := utils.CleanCreateFolder(folder)
		if err != nil {
			return nil, err
		}
	} else {
		//TODO Manage log for existing folder, reactivate
	}
	err = utils.SaveStructureByType(file, &newRepo, repositoryFormatExtension)
	if err != nil {
		return nil, err
	}
	return &newRepo, err
}

func (s *repositoryStorageManager) DeleteRepositoryByName(name string) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.Unlock()
	}()
	s.Lock()
	if !s.containsRepositoryName(name) {
		return errors.New(fmt.Sprintf("Repository name %s not present", name))
	}
	var repoName = utils.ConvertName(name)
	if repoName == defaultRepositoryName {
		return errors.New(fmt.Sprintf("Repository name %s cannot be deleted, it's the default repository", repoName))
	}
	var deleted = false
	var repoRefs = make([]RepositoryRef, 0)
	for _, repo := range s.repositories.Repositories {
		if repo.Name == repoName {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repoName)
			err := utils.DeleteFileOrFolder(folder)
			if err != nil {
				//TODO: Manage logging of error
				return err
			} else {
				deleted = true
			}
		} else {
			repoRefs = append(repoRefs, repo)
		}
	}
	if deleted {
		s.repositories.Repositories = repoRefs
		return s.SavePoint()
	} else {
		err = errors.New(fmt.Sprintf("Repository %s not found in list", name))
	}
	return err
}

func (s *repositoryStorageManager) DeleteRepositoryById(id string) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.Unlock()
	}()
	s.Lock()
	if !s.containsRepositoryId(id) {
		return errors.New(fmt.Sprintf("Repository id %s not present", id))
	}
	if id == s.getDefefaultRepositoryId() {
		return errors.New(fmt.Sprintf("Repository id %s cannot be deleted, it's the default repository", id))
	}
	var deleted = false
	var repoRefs = make([]RepositoryRef, 0)
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name)
			err := utils.DeleteFileOrFolder(folder)
			if err != nil {
				//TODO: Manage logging of error
				return err
			} else {
				deleted = true
			}
		} else {
			repoRefs = append(repoRefs, repo)
		}
	}
	if deleted {
		s.repositories.Repositories = repoRefs
		return s.SavePoint()
	} else {
		err = errors.New(fmt.Sprintf("Repository %s not found in list", id))
	}
	return err
}

func (s *repositoryStorageManager) RenameRepository(oldName string, newName string) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.Unlock()
	}()
	s.Lock()
	if !s.containsRepositoryName(oldName) {
		return errors.New(fmt.Sprintf("Repository name %s not present", oldName))
	}
	var repoName = utils.ConvertName(oldName)
	var newRepoName = utils.ConvertName(newName)
	if repoName == defaultRepositoryName {
		return errors.New(fmt.Sprintf("Repository name %s cannot be renamed, it's the default repository", repoName))
	}
	var renamed = false
	for ind, repo := range s.repositories.Repositories {
		if repo.Name == repoName {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repoName)
			var newFolder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, newRepoName)
			err := os.Rename(folder, newFolder)
			if err != nil {
				//TODO: Manage logging of error
				return err
			} else {
				s.repositories.Repositories[ind].Name = newRepoName
				renamed = true
				break
			}
		}
	}
	if renamed {
		return s.SavePoint()
	} else {
		err = errors.New(fmt.Sprintf("Repository %s not found in list", oldName))
	}
	return err
}

func (s *repositoryStorageManager) ListRepositoryCharts(id string) ([]model.Chart, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.RUnlock()
	}()
	s.RLock()
	var outList = make([]model.Chart, 0)
	if !s.containsRepositoryId(id) {
		return outList, errors.New(fmt.Sprintf("Repository id %s not present", id))
	}
	var found = false
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name)
			err := utils.DeleteFileOrFolder(folder)
			if err != nil {
				//TODO: Manage logging of error
				return outList, err
			} else {
				found = true
				var file = fmt.Sprintf(repositoryDetailsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, repositoryFormatExtension)
				if fs, err := os.Stat(file); err == nil {
					if fs.IsDir() {
						//TODO: Manage logs of folder and not file case
						return outList, errors.New(fmt.Sprintf("No repository id: %s contains folder instead file path: %s", id, file))
					} else {
						ext, err := utils.GetPathExtension(file)
						if err == nil {
							//TODO: Manage logs of folder and not file extension
							return outList, err
						}
						var extens = utils.FormatType(strings.ToLower(ext))
						if extens != repositoryFormatExtension {
							//TODO: Manage logs of file wrong etension
							return outList, errors.New(fmt.Sprintf("No repository id: %s contains file path: %s with wrong extension, expected: %v", id, file, repositoryFormatExtension))
						}
						var repo = model.Repository{}
						err = utils.LoadStructureByType(file, &repo, repositoryFormatExtension)
						if err == nil {
							//TODO: Manage logs of folder and not file extension
							return outList, err
						}
						outList = append(outList, repo.Charts...)
						break
					}
				} else {
					//TODO: Manage logging of no file read
					return outList, errors.New(fmt.Sprintf("No repository id: %s doesn't contain any chart index file: %s", id, file))
				}
			}
		}
	}
	if !found {
		return outList, errors.New(fmt.Sprintf("No repository id: %s found in the list", id))
	}
	return outList, err
}

func (s *repositoryStorageManager) ListRepositoryKubeFiles(id string) ([]model.KubeFile, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.RUnlock()
	}()
	s.RLock()
	var outList = make([]model.KubeFile, 0)
	if !s.containsRepositoryId(id) {
		return outList, errors.New(fmt.Sprintf("Repository id %s not present", id))
	}
	var found = false
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name)
			err := utils.DeleteFileOrFolder(folder)
			if err != nil {
				//TODO: Manage logging of error
				return outList, err
			} else {
				found = true
				var file = fmt.Sprintf(repositoryDetailsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, repositoryFormatExtension)
				if fs, err := os.Stat(file); err == nil {
					if fs.IsDir() {
						//TODO: Manage logs of folder and not file case
						return outList, errors.New(fmt.Sprintf("No repository id: %s contains folder instead file path: %s", id, file))
					} else {
						ext, err := utils.GetPathExtension(file)
						if err == nil {
							//TODO: Manage logs of folder and not file extension
							return outList, err
						}
						var extens = utils.FormatType(strings.ToLower(ext))
						if extens != repositoryFormatExtension {
							//TODO: Manage logs of file wrong etension
							return outList, errors.New(fmt.Sprintf("No repository id: %s contains file path: %s with wrong extension, expected: %v", id, file, repositoryFormatExtension))
						}
						var repo = model.Repository{}
						err = utils.LoadStructureByType(file, &repo, repositoryFormatExtension)
						if err == nil {
							//TODO: Manage logs of folder and not file extension
							return outList, err
						}
						outList = append(outList, repo.KubeFiles...)
						break
					}
				} else {
					//TODO: Manage logging of no file read
					return outList, errors.New(fmt.Sprintf("No repository id: %s doesn't contain any chart index file: %s", id, file))
				}
			}
		}
	}
	if !found {
		return outList, errors.New(fmt.Sprintf("No repository id: %s found in the list", id))
	}
	return outList, err
}

func (s *repositoryStorageManager) BackupRepository(id string, archiveFile string, useZipFormat bool) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.RUnlock()
	}()
	s.RLock()
	if !s.containsRepositoryId(id) {
		return errors.New(fmt.Sprintf("Repository id %s not present", id))
	}
	var found = false
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name)

			if fs, errS := os.Stat(folder); err == nil {
				found = true
				if !fs.IsDir() {
					//TODO: Manage logging of error
					return errors.New(fmt.Sprintf("File: %s is not regular folder", folder))
				}
				if useZipFormat {
					//TODO: Logging backup repo by zip format
					err = utils.ZipCompress(folder, archiveFile)
				} else {
					//TODO: Logging backup repo by tar format
					err = utils.TarCompress(folder, archiveFile, true)
				}
				break
			} else {
				//TODO: Manage logging of error
				return errS
			}
		}
	}
	if !found {
		return errors.New(fmt.Sprintf("No repository id: %s found in the list", id))
	}
	return err
}

func (s *repositoryStorageManager) RestoreRepository(archiveFile string, useZipFormat bool, forceCreate bool) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.Unlock()
	}()
	s.Lock()
	if fs, err := os.Stat(archiveFile); err == nil {
		if fs.IsDir() {
			return errors.New(fmt.Sprintf("Archive %s is folder, and not regular file!!", archiveFile))
		}
	} else {
		return errors.New(fmt.Sprintf("Archive file %s doesn't exist!!", archiveFile))
	}
	var uncompressed = false
	var tmpFolder = utils.GetTempFolder(utils.GetRandPath())
	err = utils.CleanCreateFolder(tmpFolder)
	if err != nil {
		//TODO: Logging unable to create temp folder
		return err
	}
	uncompressed = true
	if useZipFormat {
		//TODO: Logging restore repo by zip format
		err = utils.ZipUnCompress(archiveFile, tmpFolder)
	} else {
		//TODO: Logging restore repo by tar format
		err = utils.TarUnCompress(archiveFile, tmpFolder, true)
	}
	if err != nil {
		//TODO: Logging unable to uncompress archive file
		return err
	}

	if !uncompressed {
		return errors.New(fmt.Sprintf("No repository id: %s uncompressed in the list", id))
	}
	return err
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
	} else {
		return s, s.createDefault()
	}
	return s, nil
}

func (s *repositoryStorageManager) createDefault() error {
	_, err := s.CreateRepository(defaultRepositoryName)
	if err != nil {
		return err
	}
	return s.Refresh()
}

func (s *repositoryStorageManager) SavePoint() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.Unlock()
	}()
	s.Lock()
	var file = fmt.Sprintf(repositoryIndeTemplate, s.dataFolder, os.PathSeparator, repositoryFormatExtension)
	if utils.ExistsFileOrFolder(file) {
		_ = utils.DeleteFileOrFolder(file)
	}
	err = utils.SaveStructureByType(file, s.repositories, repositoryFormatExtension)
	return err
}

func (s *repositoryStorageManager) Refresh() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.RUnlock()
	}()
	s.RLock()
	var file = fmt.Sprintf(repositoryIndeTemplate, s.dataFolder, os.PathSeparator, repositoryFormatExtension)
	err = utils.LoadStructureByType(file, s.repositories, repositoryFormatExtension)
	return err
}

var repositoryStorageManagerSingleton model.RepositoryStorageManager

func GetRepositoryStorageManagerSingleton(dataFolder string) (model.RepositoryStorageManager, error) {
	var err error
	if repositoryStorageManagerSingleton == nil {
		repositoryStorageManagerSingleton, err = (&repositoryStorageManager{
			dataFolder: dataFolder,
			repositories: &Repositories{
				Repositories: make([]RepositoryRef, 0),
				Created:      time.Now(),
				Updated:      time.Now(),
			},
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
		repositories: &Repositories{
			Repositories: make([]RepositoryRef, 0),
			Created:      time.Now(),
			Updated:      time.Now(),
		},
	}).Initialize()

}
