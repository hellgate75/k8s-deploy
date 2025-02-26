package integration

import (
	"errors"
	"fmt"
	"github.com/hellgate75/k8s-deploy/utils"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func checkPath() {
	d, err := utils.GetExecutionDir()
	if err != nil {
		d, _ = os.Getwd()
	}
	var lst = utils.GetOSPathList()
	if ! utils.StringsListContainItem(d, lst, false) {
		utils.AddToOSPathList(d)
	}
}

func checkPresenctOfHelm() bool {
	out, err := utils.ExecuteCommand("helm", "--help")
	return err == nil && len(out) > 0
}

func checkPresenctOfKubectl() bool {
	out, err := utils.ExecuteCommand("kubectl", "--help")
	return err == nil && len(out) > 0
}

var helmVerTextInit="<a href=\"/helm/helm/releases/tag/"

func getHelmLatestVersion() string {
	resp, err := http.Get("https://github.com/helm/helm/releases/latest")
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return "v3.2.1"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(data) == 0 {
		return "v3.2.1"
	}
	text := string(data)
	val1 := strings.Split(text, helmVerTextInit)[1]
	ver := strings.Split(val1, "\"")[0]
	return ver
}


func getKubectlLatestVersion() string {
	resp, err := http.Get("https://storage.googleapis.com/kubernetes-release/release/stable.txt")
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return "v1.18.2"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(data) == 0 {
		return "v1.18.2"
	}
	return strings.Split(string(data), "\n")[0]
}

func downloadInstallHelm() error {
	fmt.Printf("Installing Helm ...\n")
	version := getHelmLatestVersion()
	fmt.Printf("Latest version is: %s\n", version)
	dir := os.TempDir()
	file := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, "helm")
	var url string
	var extension = "zip"
	if runtime.GOOS == "windows" {
		file = file + ".zip"
		url = fmt.Sprintf("https://get.helm.sh/helm-%s-windows-amd64.zip", version)
	} else if runtime.GOOS == "linux" {
		extension="tar.gz"
		file = file + ".tar.gz"
		url = fmt.Sprintf("https://get.helm.sh/helm-%s-linux-arm64.tar.gz", version)
	} else if runtime.GOOS == "darwin" {
		extension="tar.gz"
		file = file + ".tar.gz"
		url = fmt.Sprintf("https://get.helm.sh/helm-%s-darwin-amd64.tar.gz", version)
	} else {
		return errors.New("Unimplemented os:" + runtime.GOOS)
	}
	fmt.Printf("Saving download in: %s\n", file)
	fmt.Printf("Download url: %s\n", url)
	err := downloadFileTo(url, file)
	if err != nil {
		fmt.Printf("Errors downloading Helm: %v\n", err)
		return err
	}
	fmt.Printf("Move file: %s to binary folder\n", file)
	archivePath, err := moveToBinaryFolder(file)
	fmt.Printf("Uncompressing archive: %s...\n", archivePath)
	err = uncompressFile(archivePath, extension == "zip", "helm")
	fmt.Printf("Delete archive in path: %s\n", archivePath)
	if err == nil {
		err = utils.DeleteFileOrFolder(archivePath)
	} else {
		errX := utils.DeleteFileOrFolder(archivePath)
		fmt.Printf("Error removing emergency the file %s -> Error: %v", archivePath, errX)
	}
	return err
}


func uncompressFile(file string, zip bool, filter string) error {
	d, err := utils.GetExecutionDir()
	if err != nil {
		d, _ = os.Getwd()
	}
	if zip {
		err = utils.ZipUnCompressFilter(file, d, filter)
	} else {
		err = utils.TarUnCompressFilter(file, d, true, filter)
	}
	return nil
}

func moveToBinaryFolder(file string) (string, error) {
	d, err := utils.GetExecutionDir()
	if err != nil {
		d, _ = os.Getwd()
	}
	fmt.Printf("Moving file: %s to folder: %s\n", file, d)
	_, path, err := utils.MoveFileToFolder(file, d)
	return path, err
	//getHome
}

func downloadInstallKubectl() error {
	fmt.Printf("Installing Kubectl ...\n")
	version := getKubectlLatestVersion()
	fmt.Printf("Latest version is: %s\n", version)
	dir := os.TempDir()
	file := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, "kubectl")
	var url string
	if runtime.GOOS == "windows" {
		file = file + ".exe"
		url = fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/windows/amd64/kubectl.exe", version)
	} else if runtime.GOOS == "linux" {
		url = fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/linux/amd64/kubectl", version)
	} else if runtime.GOOS == "darwin" {
		url = fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/darwin/amd64/kubectl", version)
	} else {
		return errors.New("Unimplemented os:" + runtime.GOOS)
	}
	fmt.Printf("Saving download in: %s\n", file)
	fmt.Printf("Download url: %s\n", url)
	err := downloadFileTo(url, file)
	if err != nil {
		fmt.Printf("Errors downloading Kubectl: %v\n", err)
		return err
	}
	_, err = moveToBinaryFolder(file)
	return err
}

func downloadFileTo(url string, file string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Invalid status code: %v", resp.StatusCode))
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, d, 0777)
	return err
}