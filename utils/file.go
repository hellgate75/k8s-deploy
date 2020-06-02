package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Create a folder if it doesn't exist
func CreateFolder(folder string) error {
	if _, err := os.Stat(folder); err != nil {
		return os.MkdirAll(folder, 0666)
	}
	return nil
}

// Create a folder and if it exists, remove it before creation
func CleanCreateFolder(fileOrFolder string) error {
	if st, err := os.Stat(fileOrFolder); err == nil {
		if st.IsDir() {
			err = os.RemoveAll(fileOrFolder)
		} else {
			err = os.Remove(fileOrFolder)
		}
		if err != nil {
			return err
		}
	}
	return os.MkdirAll(fileOrFolder, 0666)
}

// Gets a file/folder path extension
func GetPathExtension(fileOrFolder string) (string, error) {
	if "" == fileOrFolder || !strings.Contains(fileOrFolder, ".") {
		return "", errors.New(fmt.Sprintf("Unable to determine extension for file or folder: %s", fileOrFolder))
	}
	tkns := strings.Split(fileOrFolder, ".")
	return tkns[len(tkns)-1], nil
}

// Delete a folder if it exists
func DeleteFileOrFolder(fileOrFolder string) error {
	if st, err := os.Stat(fileOrFolder); err == nil {
		_ = os.Chmod(fileOrFolder, 660)
		if st.IsDir() {
			fmt.Printf("Deleting folder %s...\n", fileOrFolder)
			return os.RemoveAll(fileOrFolder)
		}
		fmt.Printf("Deleting file %s...\n", fileOrFolder)
		return os.Remove(fileOrFolder)
	} else {
		fmt.Printf("File or folder %s doean't exist!!", fileOrFolder)
	}
	return nil
}

// Check if a folder exists
func ExistsFileOrFolder(fileOrFolder string) bool {
	_, err := os.Stat(fileOrFolder)
	return err == nil
}

// Copy file into a folder and return number of copied bytes, the destination file path and eventually copy operation error
func CopyFileToFolder(filePath string, folder string) (int64, string, error) {
	if !ExistsFileOrFolder(folder) {
		err := CreateFolder(folder)
		if err != nil {
			return 0, "", err
		}
	}
	if !ExistsFileOrFolder(filePath) {
		return 0, "", errors.New(fmt.Sprintf("File %s doesn't exist, cannot copy to folder %s", filePath, folder))
	}
	st, _ := os.Stat(filePath)
	_, fileName := filepath.Split(filePath)
	var dstFilePath = filepath.Join(folder, fileName)
	if ExistsFileOrFolder(dstFilePath) {

	}
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return 0, "", err
	}
	defer dstFile.Close()
	srcFile, err := os.Open(filePath)
	if err != nil {
		return 0, "", err
	}
	defer srcFile.Close()
	n, err := io.CopyN(dstFile, srcFile, st.Size())
	return n, dstFilePath, err
}

// Move file into a folder and return number of moved file bytes, the destination file path and eventually move operation error
func MoveFileToFolder(filePath string, folder string) (int64, string, error) {
	n, path, err := CopyFileToFolder(filePath, folder)
	if err != nil {
		return n, path, err
	}
	return n, path, DeleteFileOrFolder(filePath)
}

// Get Binaries execution folder
func GetExecutionDir() (string, error) {
	exec, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exec), nil
}

// Get User Home execution folder
func GetUserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

func GetOSPathList() []string {
	path := os.Getenv("Path")
	if path == "" {
		path = os.Getenv("PATH")
	}
	if path == "" {
		return make([]string, 0)
	}
	return strings.Split(path, fmt.Sprintf("%c", os.PathListSeparator))
}

func AddToOSPathList(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		return errors.New(fmt.Sprintf("Folder %s doesn't exist", dir))
	}

	if path := os.Getenv("Path"); path != "" {
		path = fmt.Sprintf("%s%c%s", path, os.PathListSeparator, dir)
		os.Setenv("Path", path)
	} else if path := os.Getenv("PATH"); path != "" {
		path = fmt.Sprintf("%s%c%s", path, os.PathListSeparator, dir)
		os.Setenv("PATH", path)
	} else {
		os.Setenv("PATH", path)
	}
	return nil
}

func GetRandPath() string {
	return uuid.New().String()
}

func GetTempFolder(folder string) string {
	return fmt.Sprintf("%s%c%s", os.TempDir(), os.PathSeparator, folder)
}
