// Copyright 2020 Re-Bind Author (Fabrizio Torelli). All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package rest

import (
	"github.com/hellgate75/k8s-deploy/model"
)

const (
	DefaultDatabaseNamePrefix  		string = "default-k8s-deploy"
	DefaultRepositoryStorageFolder  string = "/var/k8s-deploy/repository"
	DefaultSchedulerStorageFolder   string = "/var/k8s-deploy/scheduler"
	DefaultExecutorStorageFolder    string = "/var/k8s-deploy/scheduler"
	DefaultConfigFolder             string = "/etc/k8s-deploy"
	DefaultLogFileFolder            string = "/var/log/k8s-deploy"
	DefaultLogFileLevel             string = "DEBUG"
	DefaultRepositoryRestServerPort int    = 8089
	DefaultSchedulerRestServerPort	 int    = 8090
	DefaultExecutorRestServerPort   int    = 8091
	DefaultIpAddress                       = "0.0.0.0"
)

type UpdateListItem struct {
	Name  string `yaml:"name" json:"name" xml:"name"`
	Value string `yaml:"value" json:"value" xml:"value"`
	Index int    `yaml:"index" json:"index" xml:"index"`
}


type UpdateItem struct {
	Name  string `yaml:"name" json:"name" xml:"name"`
	Value string `yaml:"value" json:"value" xml:"value"`
}

type UpdateRequest struct {
	ListData 	[]UpdateListItem 	`yaml:"listItems,omitempty" json:"listItems,omitempty" xml:"list-item,omitempty"`
	RecordData 	[]UpdateItem 		`yaml:"items,omitempty" json:"items,omitempty" xml:"item,omitempty"`
	Request 	model.Request  		`yaml:"request" json:"request" xml:"request"`
}

type DeleteListItem struct {
	Name  string `yaml:"name" json:"name" xml:"name"`
	Index int    `yaml:"index" json:"index" xml:"index"`
}

type DeleteItem struct {
	Name  string `yaml:"name" json:"name" xml:"name"`
}

type DeleteRequest struct {
	ListData 	[]DeleteListItem 	`yaml:"listItems,omitempty" json:"listItems,omitempty" xml:"list-item,omitempty"`
	RecordData 	[]DeleteItem 		`yaml:"items,omitempty" json:"items,omitempty" xml:"item,omitempty"`
	Request 	model.Request  		`yaml:"request" json:"request" xml:"request"`
}

type TemplateDataType struct {
	Method  string      `yaml:"method" json:"method" xml:"method"`
	Header  []string    `yaml:"header" json:"header" xml:"header"`
	Query   []string    `yaml:"query" json:"query" xml:"query"`
	Request interface{} `yaml:"request" json:"request" xml:"request"`
}
type TemplateResponse struct {
	Templates []TemplateDataType `yaml:"template" json:"template" xml:"template"`
}

