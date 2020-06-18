package model

import "time"

type RepositoryRef struct {
	Id   string
	Name string
}

type Repositories struct {
	Repositories []RepositoryRef `yaml:"repositories" json:"repositories" xml:"repository"`
	Created      time.Time       `yaml:"created" json:"created" xml:"created"`
	Updated      time.Time       `yaml:"updated" json:"updated" xml:"updated"`
}

type ChartList struct {
	RepoName string      `yaml:"repository" json:"repository" xml:"repository"`
	Charts   []ChartInfo `yaml:"charts" json:"charts" xml:"chart"`
}

type KubernetesFileList struct {
	RepoName string               `yaml:"repository" json:"repository" xml:"repository"`
	Files    []KubernetesFileInfo `yaml:"files" json:"files" xml:"file"`
}
