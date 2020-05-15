package model

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ArgumentsList []string

func (i *ArgumentsList) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *ArgumentsList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *ArgumentsList) Get(index int) string {
	if index > 0 && index < len(*i) {
		return (*i)[index]
	}
	return ""
}

type KubeRepoConfig struct {
	DataDirPath         string `yaml:"dataDir" json:"dataDir" xml:"data-dir"`
	ConfigDirPath       string `yaml:"configDir" json:"configDir" xml:"config-dir"`
	ListenIP            string `yaml:"listenIp" json:"listenIp" xml:"listen-ip"`
	ListenPort          int    `yaml:"listenPort" json:"listenPort" xml:"listen-port"`
	TlsCert             string `yaml:"tlsCertFilePath" json:"tlsCertFilePath" xml:"tls-cert-file-path"`
	TlsKey              string `yaml:"tlsKeyFilePath" json:"tlsKeyFilePath" xml:"tls-key-file-path"`
	EnableFileLogging   bool   `yaml:"enableFileLogging" json:"enableFileLogging" xml:"enable-file-logging"`
	LogVerbosity        string `yaml:"logVerbosity" json:"logVerbosity" xml:"log-verbosity"`
	LogFilePath         string `yaml:"logFilePath" json:"logFilePath" xml:"log-file-path"`
	EnableLogRotate     bool   `yaml:"enableLogRotate" json:"enableLogRotate" xml:"enable-log-rotate"`
	LogMaxFileSize      int64  `yaml:"logMaxFileSize" json:"logMaxFileSize" xml:"log-max-file-size"`
	LogFileCount        int    `yaml:"logFileCount" json:"logFileCount" xml:"log-file-count"`
	MongoDbEnabled 		bool	`yaml:"mongoDbEnabled" json:"mongoDbEnabled" xml:"mongo-db-enabled"`
	MongoDbHost 		string	`yaml:"mongoDbHost" json:"mongoDbHost" xml:"mongo-db-host"`
	MongoDbPort 		int		`yaml:"mongoDbPort" json:"mongoDbPort" xml:"mongo-db-port"`
	MongoDbUser 		string	`yaml:"mongoDbUser" json:"mongoDbUser" xml:"mongo-db-user"`
	MongoDbPassword 	string	`yaml:"mongoDbPassword" json:"mongoDbPassword" xml:"mongo-db-password"`
}

func (conf KubeRepoConfig) ToJson() string {
	bArr, err := json.Marshal(conf)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(bArr)
}

func (conf KubeRepoConfig) ToYaml() string {
	bArr, err := yaml.Marshal(conf)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(bArr)
}

func (conf KubeRepoConfig) ToXml() string {
	bArr, err := xml.Marshal(conf)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(bArr)
}

func SaveConfig(path string, name string, config interface{}) error {
	if _, err := os.Stat(path); err != nil {
		_ = os.MkdirAll(path, 0660)
	}
	fileFullPath := fmt.Sprintf("%s%c%s.yaml", path, os.PathSeparator, name)
	bArr, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileFullPath, bArr, 0666)
	return err
}

func LoadConfig(path string, name string, config interface{}) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	fileFullPath := fmt.Sprintf("%s%c%s.yaml", path, os.PathSeparator, name)
	if _, err := os.Stat(fileFullPath); err != nil {
		return err
	}
	bArr, err := ioutil.ReadFile(fileFullPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bArr, config)
	return err
}
