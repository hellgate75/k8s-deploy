package model

import "github.com/hellgate75/k8s-deploy/model"

func RemoveChartsDuplicates(c []model.Chart) []model.Chart {
	var charts = make([]model.Chart, 0)
	var reg = make(map[string]bool)
	for _, ch := range c {
		if _, ok := reg[ch.Name]; !ok {
			reg[ch.Name] = true
			charts = append(charts, ch)
		}
	}
	return charts
}

func RemoveKubernetesFilesDuplicates(c []model.KubernetesFile) []model.KubernetesFile {
	var kubernetesFiles = make([]model.KubernetesFile, 0)
	var reg = make(map[string]bool)
	for _, kf := range c {
		if _, ok := reg[kf.Name]; !ok {
			reg[kf.Name] = true
			kubernetesFiles = append(kubernetesFiles, kf)
		}
	}
	return kubernetesFiles
}

func ListContainerChart(charts []model.Chart, c model.Chart) bool {
	for _, ch := range charts {
		if ch.Name == c.Name {
			return true
		}
	}
	return false
}

func ListContainerKubernetesFile(kubernetesFiles []model.KubernetesFile, kf model.KubernetesFile) bool {
	for _, skf := range kubernetesFiles {
		if skf.Name == kf.Name {
			return true
		}
	}
	return false
}
