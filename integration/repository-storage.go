package integration

import (
	"errors"
	"fmt"
	"github.com/hellgate75/k8s-deploy/log"
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
	umodel "github.com/hellgate75/k8s-deploy/utils/model"
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

type repositoryStorageManager struct {
	sync.RWMutex
	dataFolder   string
	repositories *model.Repositories
	logger       log.Logger
}

func (s *repositoryStorageManager) GetRepositoryList() []model.Repository {
	s.logger.Infof("RepositoryStorageManager::GetRepositoryList() ...")
	var out = make([]model.Repository, 0)
	for _, ref := range s.repositories.Repositories {
		repo, err := s.GetRepositoryById(ref.Id)
		if err != nil {
			if s.logger != nil {
				s.logger.Errorf("RepositoryStorageManager::GetRepositoryList() - Repository id: %s - Error: %v", ref.Id, err)
			}
		} else {
			out = append(out, *repo)
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
			return s.GetRepositoryById(repo.Id)
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
			// Load repository data
			var file = fmt.Sprintf(repositoryDetailsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, repositoryFormatExtension)
			var repository = model.Repository{}
			err := utils.LoadStructureByType(file, &repository, repositoryFormatExtension)
			if err != nil {
				return nil, err
			}
			// Load charts related to the repository
			file = fmt.Sprintf(repositoryChartsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
			var chartsFileList = model.ChartList{
				RepoName: repo.Name,
				Charts:   make([]model.Chart, 0),
			}
			err = utils.LoadStructureByType(file, &chartsFileList, repositoryFormatExtension)
			if err != nil {
				return nil, err
			}
			//Sets the discovered charts in the repository structure
			repository.ReplaceCharts(chartsFileList.Charts...)
			// Load kubernetes files related to the repository
			file = fmt.Sprintf(repositoryKubefilesIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
			var kubernetesFileList = model.KubeFileList{
				RepoName: repo.Name,
				Files:    make([]model.KubernetesFile, 0),
			}
			err = utils.LoadStructureByType(file, &kubernetesFileList, repositoryFormatExtension)
			if err != nil {
				return nil, err
			}
			//Sets the discovered kubernetes files in the repository structure
			repository.ReplaceKubernetesFiles(kubernetesFileList.Files...)
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
	}()
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("Repository name cannot be empty or without significant digits or letters")
	}
	if s.containsRepositoryName(name) {
		return nil, errors.New(fmt.Sprintf("Repository name %s already present", name))
	}
	var repoName = utils.ConvertName(name)
	var newRepo = model.CreateRepository(
		utils.NewUID(),
		repoName,
		model.StateCreated)
	// Create Repository Files
	err = s.saveRepository(repoName, newRepo)
	if err != nil {
		return nil, err
	}

	// Create Repository Charts File
	err = s.saveCharts(repoName, newRepo.GetCharts())
	if err != nil {
		return nil, err
	}

	// Create Repository Kubernetes Charts
	err = s.saveKubernetesFiles(repoName, newRepo.GetKubernetesFiles())
	if err != nil {
		return nil, err
	}

	rr := model.RepositoryRef{
		Id:   newRepo.Id,
		Name: newRepo.Name,
	}
	s.repositories.Repositories = append(s.repositories.Repositories, rr)
	err = s.SavePoint()
	if err != nil {
		return nil, err
	}
	return &newRepo, err
}

func (s *repositoryStorageManager) UpdateRepository(id string, r model.Repository) (*model.Repository, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	if len(r.Name) == 0 {
		return nil, errors.New(fmt.Sprintf("Cannot update any repository with id: %s without a repoName", id))
	}
	repo, err := s.GetRepositoryById(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot find any repository with id: %s, error: %v", id, err))
	}
	if repo == nil {
		return nil, errors.New(fmt.Sprintf("Cannot find any repository with id: %s it was not found", id))
	}
	var repoCharts = repo.GetCharts()
	var repoKubeFiles = repo.GetKubernetesFiles()
	repoCharts = append(repoCharts, r.GetCharts()...)
	repoKubeFiles = append(repoKubeFiles, r.GetKubernetesFiles()...)
	r.Name = utils.ConvertName(r.Name)
	repo.ReplaceCharts(umodel.RemoveChartsDuplicates(repoCharts)...)
	repo.ReplaceKubernetesFiles(umodel.RemoveKubernetesFilesDuplicates(repoKubeFiles)...)
	//	if repo.Name == r.Name {
	//		return nil, errors.New(fmt.Sprintf("Cannot rename repository with same name for repository id: %s it was not found", id))
	//	}
	var oldName = repo.Name
	if oldName != repo.Name {
		repo.Name = r.Name
	}
	repoByName, err := s.GetRepository(r.Name)
	if err == nil && repoByName.Id != repo.Id {
		return nil, errors.New(fmt.Sprintf("Repository name %s already exist with anther id: %s", repoByName.Name, repoByName.Id))
	}
	var oldFolder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, oldName)
	var newFolder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, r.Name)
	if !utils.ExistsFileOrFolder(oldFolder) {
		if s.logger != nil {
			s.logger.Debugf("Creating new folder for repository: %s", repo.Name)
		}
		err := utils.CleanCreateFolder(newFolder)
		if err != nil {
			return nil, err
		}
	} else {
		if oldName != repo.Name {
			if s.logger != nil {
				s.logger.Debugf("Renaming folder for repository: %s to repository %s", oldName, repo.Name)
			}
			err = os.Rename(oldFolder, newFolder)
			if err != nil {
				return nil, err
			}
		} else {
			if s.logger != nil {
				s.logger.Debugf("No name changes in repository: %s", repo.Name)
			}
		}
	}
	err = s.saveRepository(repo.Name, *repo)
	if err != nil {
		return nil, err
	}

	if oldName != repo.Name {
		var changed = false
		for idx, rr := range s.repositories.Repositories {
			if rr.Id == repo.Id {
				s.repositories.Repositories[idx].Name = r.Name
				changed = true
			}
		}
		if !changed {
			return nil, errors.New(fmt.Sprintf("Unable to continue, no changes because id: %s was not found in cluster list.", id))
		}
		err = s.SavePoint()
	}
	err = s.saveCharts(repo.Name, repo.GetCharts())
	if err != nil {
		return nil, err
	}
	err = s.saveKubernetesFiles(repo.Name, repo.GetKubernetesFiles())
	if err != nil {
		return nil, err
	}
	return repo, err
}

func (s *repositoryStorageManager) OverrideRepository(id string, r model.Repository) (*model.Repository, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	if len(r.Name) == 0 {
		return nil, errors.New(fmt.Sprintf("Cannot override any repository with id: %s without a repoName", id))
	}
	repo, err := s.GetRepositoryById(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot find any repository with id: %s, error: %v", id, err))
	}
	if repo == nil {
		return nil, errors.New(fmt.Sprintf("Cannot find any repository with id: %s it was not found", id))
	}
	r.Name = utils.ConvertName(r.Name)
	var oldName = repo.Name
	var newName = r.Name
	var changeName = false
	if oldName != repo.Name {
		changeName = true
	}
	var mergeWithRepo = false
	mergeRepository, err := s.GetRepository(r.Name)
	if err == nil && mergeRepository.Id != repo.Id {
		if s.logger != nil {
			s.logger.Warnf(fmt.Sprintf("Repository name %s already exist with anther id: %s, merging the repositories", mergeRepository.Name, mergeRepository.Id))
		}
		mergeWithRepo = true
	}
	var file = fmt.Sprintf(repositoryDetailsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, r.Name, os.PathSeparator, repositoryFormatExtension)
	if changeName {
		var oldFolder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, oldName)
		var newFolder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, newName)
		if mergeWithRepo {
			var mergeFolder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, mergeRepository.Name)
			if !utils.ExistsFileOrFolder(oldFolder) && !utils.ExistsFileOrFolder(mergeFolder) {
				if s.logger != nil {
					s.logger.Debugf("Creating new folder for repository: %s", newName)
				}
				err := utils.CleanCreateFolder(newFolder)
				if err != nil {
					return nil, err
				}
			} else {
				if changeName {
					// We have to move the files, only the listed repository files
					// From old repo to merge repo

					if s.logger != nil {
						s.logger.Debugf("Renaming folder for repository: %s to repository %s", oldName, repo.Name)
					}
					err = os.Rename(oldFolder, newFolder)
					if err != nil {
						return nil, err
					}
				} else {
					// We have to move the files, only the listed repository files
					// From repo to merge repo
					if s.logger != nil {
						s.logger.Debugf("No name changes in repository: %s", repo.Name)
					}
				}
			}
		} else {
			if !utils.ExistsFileOrFolder(oldFolder) {
				if s.logger != nil {
					s.logger.Debugf("Creating new folder for repository: %s", repo.Name)
				}
				err := utils.CleanCreateFolder(newFolder)
				if err != nil {
					return nil, err
				}
			} else {
				if oldName != repo.Name {
					if s.logger != nil {
						s.logger.Debugf("Renaming folder for repository: %s to repository %s", oldName, repo.Name)
					}
					err = os.Rename(oldFolder, newFolder)
					if err != nil {
						return nil, err
					}
				} else {
					if s.logger != nil {
						s.logger.Debugf("No name changes in repository: %s", repo.Name)
					}
				}
			}
		}
	} else {

	}
	err = utils.SaveStructureByType(file, repo, repositoryFormatExtension)
	if err != nil {
		return nil, err
	}
	if oldName != repo.Name {
		var changed = false
		for idx, rr := range s.repositories.Repositories {
			if rr.Id == repo.Id {
				s.repositories.Repositories[idx].Name = r.Name
				changed = true
			}
		}
		if !changed {
			return nil, errors.New(fmt.Sprintf("Unable to continue, no changes because id: %s was not found in cluster list.", id))
		}
		err = s.SavePoint()
	}
	return repo, err
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
	var repoRefs = make([]model.RepositoryRef, 0)
	for _, repo := range s.repositories.Repositories {
		if repo.Name == repoName {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repoName)
			err := utils.DeleteFileOrFolder(folder)
			if err != nil {
				if s.logger != nil {
					s.logger.Errorf("Error occurred deleting repository %s, Error: %v", name, err)
				}
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
	var repoRefs = make([]model.RepositoryRef, 0)
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			var folder = fmt.Sprintf(repositoryDetailsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name)
			err := utils.DeleteFileOrFolder(folder)
			if err != nil {
				if s.logger != nil {
					s.logger.Errorf("Error occurred deleting repository with id: %s, Error: %v", id, err)
				}
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
				if s.logger != nil {
					s.logger.Errorf("Error renaming repository: %s, Error: %v", repo.Name, err)
				}
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
			var folder = fmt.Sprintf(repositoryChartsFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator)
			if _, err := os.Stat(folder); err != nil {
				return outList, err
			} else {
				found = true
				var file = fmt.Sprintf(repositoryChartsIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
				if fs, err := os.Stat(file); err == nil {
					if fs.IsDir() {
						if s.logger != nil {
							s.logger.Errorf("Error recovering file %s for repository: %s, it's a folder", file, repo.Name)
						}
						return outList, errors.New(fmt.Sprintf("No repository id: %s contains folder instead file path: %s", id, file))
					} else {
						ext, err := utils.GetPathExtension(file)
						if err == nil {
							if s.logger != nil {
								s.logger.Errorf("Error recovering file %s extension for repository: %s, ", file, repo.Name)
							}
							return outList, err
						}
						var extens = utils.FormatType(strings.ToLower(ext))
						if extens != repositoryFormatExtension {
							if s.logger != nil {
								s.logger.Errorf("Error wrong file %s extension %v for repository: %s, ", file, extens, repo.Name)
							}
							return outList, errors.New(fmt.Sprintf("No repository id: %s contains file path: %s with wrong extension, expected: %v", id, file, repositoryFormatExtension))
						}
						var chartsList = model.ChartList{
							RepoName: repo.Name,
							Charts:   make([]model.Chart, 0),
						}
						err = utils.LoadStructureByType(file, &chartsList, repositoryFormatExtension)
						if err == nil {
							if s.logger != nil {
								s.logger.Errorf("Error loading file %s for repository: %s, ", file, chartsList.RepoName)
							}
							return outList, err
						}
						outList = append(outList, chartsList.Charts...)
						break
					}
				} else {
					if s.logger != nil {
						s.logger.Errorf("Error unable to loading file %s for repository: %s, ", file, repo.Name)
					}
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

func (s *repositoryStorageManager) ListRepositoryKubeFiles(id string) ([]model.KubernetesFile, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
		s.RUnlock()
	}()
	s.RLock()
	var outList = make([]model.KubernetesFile, 0)
	if !s.containsRepositoryId(id) {
		return outList, errors.New(fmt.Sprintf("Repository id %s not present", id))
	}
	var found = false
	for _, repo := range s.repositories.Repositories {
		if repo.Id == id {
			var folder = fmt.Sprintf(repositoryKubefilesFolderTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator)
			if _, err := os.Stat(folder); err != nil {
				return outList, err
			} else {
				found = true
				var file = fmt.Sprintf(repositoryKubefilesIndexTemplate, s.dataFolder, os.PathSeparator, os.PathSeparator, repo.Name, os.PathSeparator, os.PathSeparator, repositoryFormatExtension)
				if fs, err := os.Stat(file); err == nil {
					if fs.IsDir() {
						if s.logger != nil {
							s.logger.Errorf("Error recovering file %s for repository: %s, it's a folder", file, repo.Name)
						}
						return outList, errors.New(fmt.Sprintf("No repository id: %s contains folder instead file path: %s", id, file))
					} else {
						ext, err := utils.GetPathExtension(file)
						if err == nil {
							if s.logger != nil {
								s.logger.Errorf("Error recovering file %s extension for repository: %s, ", file, repo.Name)
							}
							return outList, err
						}
						var extens = utils.FormatType(strings.ToLower(ext))
						if extens != repositoryFormatExtension {
							if s.logger != nil {
								s.logger.Errorf("Error wrong file %s extension %v for repository: %s, ", file, extens, repo.Name)
							}
							return outList, errors.New(fmt.Sprintf("No repository id: %s contains file path: %s with wrong extension, expected: %v", id, file, repositoryFormatExtension))
						}
						var chartsList = model.KubeFileList{
							RepoName: repo.Name,
							Files:    make([]model.KubernetesFile, 0),
						}
						err = utils.LoadStructureByType(file, &chartsList, repositoryFormatExtension)
						if err == nil {
							if s.logger != nil {
								s.logger.Errorf("Error loading file %s for repository: %s, ", file, chartsList.RepoName)
							}
							return outList, err
						}
						outList = append(outList, chartsList.Files...)
						break
					}
				} else {
					if s.logger != nil {
						s.logger.Errorf("Error unable to loading file %s for repository: %s, ", file, repo.Name)
					}
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
					if s.logger != nil {
						s.logger.Errorf("Error required folder %s, is not a folder", folder)
					}
					return errors.New(fmt.Sprintf("File: %s is not regular folder", folder))
				}
				if useZipFormat {
					if s.logger != nil {
						s.logger.Warnf("Compressing with zip format repository %s to archive %s", repo.Name, archiveFile)
					}
					err = utils.ZipCompress(folder, archiveFile)
				} else {
					if s.logger != nil {
						s.logger.Warnf("Compressing with tar/g-zip format repository %s to archive %s", repo.Name, archiveFile)
					}
					err = utils.TarCompress(folder, archiveFile, true)
				}
				break
			} else {
				if s.logger != nil {
					s.logger.Errorf("Unable to find folder %s repository %s, Error: %s", folder, repo.Name, errS)
				}
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
		if s.logger != nil {
			s.logger.Errorf("Unable to create temporary folder for uncompressing archive %s", tmpFolder, archiveFile)
		}
		return err
	}
	uncompressed = true
	if useZipFormat {
		if s.logger != nil {
			s.logger.Warnf("Decompressing with zip format archive %s to folder %s", archiveFile, tmpFolder)
		}
		err = utils.ZipUnCompress(archiveFile, tmpFolder)
	} else {
		if s.logger != nil {
			s.logger.Warnf("Decompressing with tar-g-zip format archive %s to folder %s", archiveFile, tmpFolder)
		}
		err = utils.TarUnCompress(archiveFile, tmpFolder, true)
	}
	if err != nil {
		return err
	}

	if !uncompressed {
		return errors.New(fmt.Sprintf("Not able to restore archive: %s", archiveFile))
	}
	return err
}

func (s *repositoryStorageManager) GetRepositoryChartsManager(id string) (model.RepositoryChartManager, error) {
	r, err := s.GetRepositoryById(id)
	if err != nil {
		return nil, err
	}
	return NewRepositoryChartManager(*r, s.dataFolder, s.logger), nil
}

func (s *repositoryStorageManager) GetRepositoryKubeFilesManager(id string) (model.RepositoryKubeFilesManager, error) {
	r, err := s.GetRepositoryById(id)
	if err != nil {
		return nil, err
	}
	return NewRepositoryKubeFilesManager(*r, s.dataFolder, s.logger), nil
}

func (s *repositoryStorageManager) GetRepositoryChartsManagerByName(name string) (model.RepositoryChartManager, error) {
	r, err := s.GetRepository(name)
	if err != nil {
		return nil, err
	}
	return NewRepositoryChartManager(*r, s.dataFolder, s.logger), nil
}

func (s *repositoryStorageManager) GetRepositoryKubeFilesManagerByName(name string) (model.RepositoryKubeFilesManager, error) {
	r, err := s.GetRepository(name)
	if err != nil {
		return nil, err
	}
	return NewRepositoryKubeFilesManager(*r, s.dataFolder, s.logger), nil
}

func (s *repositoryStorageManager) Initialize() (model.RepositoryStorageManager, error) {
	s.logger.Infof("RepositoryStorageManager::Initialize() ...")
	if _, err := os.Stat(s.dataFolder); err != nil {
		err = os.MkdirAll(s.dataFolder, 666)
		if err != nil {
			return nil, err
		}
	}
	var file = fmt.Sprintf(repositoryIndeTemplate, s.dataFolder, os.PathSeparator, repositoryFormatExtension)
	if utils.ExistsFileOrFolder(file) {
		s.logger.Infof("RepositoryStorageManager::Initialize() load existing repositories ...")
		return s, s.Refresh()
	} else {
		s.logger.Infof("RepositoryStorageManager::Initialize() create new default repositories ...")
		return s, s.createDefault()
	}
	return s, nil
}

func (s *repositoryStorageManager) createDefault() error {
	_ = s.SavePoint()
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

func GetRepositoryStorageManagerSingleton(dataFolder string, logger log.Logger) (model.RepositoryStorageManager, error) {
	var err error
	if repositoryStorageManagerSingleton == nil {
		repositoryStorageManagerSingleton, err = (&repositoryStorageManager{
			dataFolder: dataFolder,
			repositories: &model.Repositories{
				Repositories: make([]model.RepositoryRef, 0),
				Created:      time.Now(),
				Updated:      time.Now(),
			},
			logger: logger,
		}).Initialize()
		if err != nil {
			repositoryStorageManagerSingleton = nil
			return nil, err
		}
	}
	return repositoryStorageManagerSingleton, err
}

func NewRepositoryStorageManager(dataFolder string, logger log.Logger) (model.RepositoryStorageManager, error) {
	return (&repositoryStorageManager{
		dataFolder: dataFolder,
		repositories: &model.Repositories{
			Repositories: make([]model.RepositoryRef, 0),
			Created:      time.Now(),
			Updated:      time.Now(),
		},
		logger: logger,
	}).Initialize()

}
