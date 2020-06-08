package model

import "github.com/hellgate75/k8s-deploy/model"

// Remove duplicates Kubernetes Helm charts in an array
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

// Remove duplicates Kubernetes files in an array
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

// Verify if a Kubernetes Helm chart array contains a single chart
func ListContainsChart(charts []model.Chart, c model.Chart) bool {
	for _, ch := range charts {
		if ch.Name == c.Name {
			return true
		}
	}
	return false
}

// Verify if a Kubernetes files array contains a single Kubernetes file
func ListContainerKubernetesFile(kubernetesFiles []model.KubernetesFile, kf model.KubernetesFile) bool {
	for _, skf := range kubernetesFiles {
		if skf.Name == kf.Name {
			return true
		}
	}
	return false
}

// Extract all matching Kubernetes Helm charts array from another list to compare, end not contained in case of invert with true value
func ExtractDiffsInChartsList(source []model.Chart, compare []model.Chart, invert bool) []model.Chart {
	var reg = make(map[string]bool)
	var out = make([]model.Chart, 0)
	for _, ch := range source {
		reg[ch.Name] = true
	}
	for _, ch := range compare {
		if _, ok := reg[ch.Name]; ok && !invert {
			out = append(out, ch)
		} else if _, ok := reg[ch.Name]; !ok {
			out = append(out, ch)
		}
	}
	return out
}

// Extract all matching Kubernetes files array from another list to compare, end not contained in case of invert with true value
func ExtractDiffsInKubernetesFilesList(source []model.KubernetesFile, compare []model.KubernetesFile, invert bool) []model.KubernetesFile {
	var reg = make(map[string]bool)
	var out = make([]model.KubernetesFile, 0)
	for _, ch := range source {
		reg[ch.Name] = true
	}
	for _, ch := range compare {
		if _, ok := reg[ch.Name]; ok && !invert {
			out = append(out, ch)
		} else if _, ok := reg[ch.Name]; !ok {
			out = append(out, ch)
		}
	}
	return out
}
