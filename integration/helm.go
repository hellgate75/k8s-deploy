package integration

import "errors"

func InitPackage() {
//	if ! checkPresenctOfHelm() {
		downloadInstallHelm()
//	}
//	if ! checkPresenctOfKubectl() {
		downloadInstallKubectl()
//	}
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