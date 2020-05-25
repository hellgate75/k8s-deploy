package integration

import (
	"errors"
	"fmt"
)

func InitPackage() {
	checkPath()
	if ! checkPresenctOfHelm() {
		err := downloadInstallHelm()
		if err != nil {
			fmt.Printf("Problems during installation of helm: %v\n", err)
		}
	}
	if ! checkPresenctOfKubectl() {
		err := downloadInstallKubectl()
		if err != nil {
			fmt.Printf("Problems during installation of kubectl: %v\n", err)
		}
	}
}

type ChartDeployRequest struct {
	ChartName		string
	ChartVersion	string
	Namespace		string
	KubeConfig		string
}

type KubeFileDeployRequest struct {
	ChartName		string
	ChartVersion	string
	Namespace		string
	KubeConfig		string
}

type HelmResponse struct {
	Code		int
	Response	string
	Output		string
	Error		error
}

func ExecuteChartRequest(req ChartDeployRequest) HelmResponse {
	return HelmResponse {
		Code: 404,
		Error: errors.New("Not Implemented!!"),
	}
}


func ExecuteKubeRequest(req KubeFileDeployRequest) HelmResponse {
	return HelmResponse {
		Code: 404,
		Error: errors.New("Not Implemented!!"),
	}
}