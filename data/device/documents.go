package device

import "github.com/hellgate75/k8s-deploy/model"

type _documentsManager struct {
	baseDataFolder	string
	repo 			*model.Repository
}

func (dm *_documentsManager) AddChart(c model.Chart)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (dm *_documentsManager) AddKubeFile(f model.KubeFile)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (dm *_documentsManager) AddChartVersion(c model.Chart, v model.Version)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) AddKubeFileVersion(f model.KubeFile, v model.Version)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) RemoveCharts(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) RemoveKubeFiles(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) RemoveChartVersions(c model.Chart, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) RemoveKubeFileVersions(f model.KubeFile, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) PurgeCharts(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) PurgeKubeFiles(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) PurgeChartVersions(c model.Chart, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) PurgeKubeFileVersions(f model.KubeFile, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) UpdateCharts(c model.Chart, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) UpdateKubeFiles(f model.KubeFile, v model.Version, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) UpdateChartVersions(c model.Chart, v model.Version, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) UpdateKubeFileVersions(f model.KubeFile, v model.Version, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) QueryCharts(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func (dm *_documentsManager) QueryKubeFiles(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) QueryChartVersions(c model.Chart, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) QueryKubeFileVersions(f model.KubeFile, q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) ListCharts()  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) ListKubeFiles()  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) ListChartVersions(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}


func (dm *_documentsManager) ListKubeFileVersions(q  ...model.Query)  model.DataResponse {
	return model.DataResponse{
		Success: false,
		Message: "Not Implemented",
	}
}

func GetDocumentDataManager(baseFolder string, repo *model.Repository) model.DocumentsDataManager {
	return &_documentsManager {
		baseDataFolder: baseFolder,
		repo: repo,
	}
}
