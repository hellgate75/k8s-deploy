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

type ChartRequest struct {
	ChartName		string
	ChartVersion	string
	Namespace		string
	KubeConfig		string
}

type KubeFileRequest struct {
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

func ExecuteChartRequest(req ChartRequest) HelmResponse {
	return HelmResponse {
		Code: 404,
		Error: errors.New("Not Implemented!!"),
	}
}


func ExecuteKubeRequest(req KubeFileRequest) HelmResponse {
	return HelmResponse {
		Code: 404,
		Error: errors.New("Not Implemented!!"),
	}
}