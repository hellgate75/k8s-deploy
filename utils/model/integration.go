package model

import (
	"github.com/hellgate75/k8s-deploy/model"
	"github.com/hellgate75/k8s-deploy/utils"
	"strconv"
	"strings"
)

// Remove duplicates Kubernetes Helm charts in an array
func RemoveChartsDuplicates(c []model.ChartInfo) []model.ChartInfo {
	var charts = make([]model.ChartInfo, 0)
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
func RemoveKubernetesFilesDuplicates(c []model.KubernetesFileInfo) []model.KubernetesFileInfo {
	var kubernetesFiles = make([]model.KubernetesFileInfo, 0)
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

type DataType byte

const (
	DataTypeString DataType = iota + 1
	DataTypeNumber
	DataTypeDecimal
	DataTypeDateTime
	DataTypeBool
)

func compareStringValues(value1, value2 string, cond model.Aggregator) bool {
	switch cond {
	case model.AggregatorEq:
		return value1 == value2
	case model.AggregatorLike:
		return strings.Contains(value1, value2)
	case model.AggregatorNotLike:
		return !strings.Contains(value1, value2)
	case model.AggregatorNeq:
		return value1 != value2
	case model.AggregatorIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNotIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return !utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNot:
		return false
	}
	return false
}

func compareNumberValues(value1, value2 string, cond model.Aggregator) bool {
	switch cond {
	case model.AggregatorEq, model.AggregatorLike:
		return value1 == value2
	case model.AggregatorNeq, model.AggregatorNotLike:
		return value1 != value2
	case model.AggregatorIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNotIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return !utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNot:
		return false
	}
	return false
}

func compareDateTimeValues(value1, value2 string, cond model.Aggregator) bool {
	switch cond {
	case model.AggregatorEq, model.AggregatorLike:
		return value1 == value2
	case model.AggregatorNeq, model.AggregatorNotLike:
		return value1 != value2
	case model.AggregatorIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNotIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return !utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNot:
		return false
	}
	return false
}

func compareBoolValues(value1, value2 string, cond model.Aggregator) bool {
	switch cond {
	case model.AggregatorEq, model.AggregatorLike:
		return value1 == value2
	case model.AggregatorNeq, model.AggregatorNotLike:
		return value1 != value2
	case model.AggregatorIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNotIn:
		// Case In
		var arr = strings.Split(value2, ",")
		return !utils.StringsListContainItem(value1, arr, true)
	case model.AggregatorNot:
		b1, err1 := strconv.ParseBool(value1)
		b2, err2 := strconv.ParseBool(value2)
		if err1 != nil || err2 != nil {
			return false
		}
		return !b1 && !b2
	}
	return false
}

func CompareValues(value1, value2 string, dType DataType, cond model.Aggregator) bool {
	switch dType {
	case DataTypeString:
		return compareStringValues(value1, value2, cond)
	case DataTypeNumber, DataTypeDecimal:
		return compareStringValues(value1, value2, cond)
	case DataTypeDateTime:
		return compareDateTimeValues(value1, value2, cond)
	case DataTypeBool:
		return compareBoolValues(value1, value2, cond)
	}
	return false
}
