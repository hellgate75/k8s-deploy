package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type FormatType string

const (
	YAML_FORMAT FormatType = "yaml"
	JSON_FORMAT FormatType = "json"
	XML_FORMAT  FormatType = "xml"
)

func LoadStructureByType(fullPath string, data interface{}, format FormatType) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var out []byte
	var f *os.File
	f, err = os.Open(fullPath)
	defer f.Close()
	if err == nil {
		out, err = ioutil.ReadAll(f)
		if err == nil {
			if format == YAML_FORMAT {
				err = yaml.Unmarshal(out, data)
			} else if format == JSON_FORMAT {
				err = json.Unmarshal(out, data)
			} else if format == XML_FORMAT {
				err = xml.Unmarshal(out, data)
			} else {
				return errors.New(fmt.Sprintf("Unable to identify given format: %v", format))
			}
		}
	}
	return err
}

func SaveStructureByType(fullPath string, data interface{}, format FormatType) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var out []byte
	if format == YAML_FORMAT {
		out, err = yaml.Marshal(data)
	} else if format == JSON_FORMAT {
		out, err = json.Marshal(data)
	} else if format == XML_FORMAT {
		out, err = xml.Marshal(data)
	} else {
		return errors.New(fmt.Sprintf("Unable to identify given format: %v", format))
	}

	if err == nil {
		err = ioutil.WriteFile(fullPath, out, 0666)
	}
	return err
}

func SaveStructureToYamlFile(folder string, filePath string, data interface{}) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var fullPath = fmt.Sprintf("%s%c%s.yaml", folder, os.PathSeparator, filePath)
	var out []byte
	out, err = yaml.Marshal(data)
	if err == nil {
		err = ioutil.WriteFile(fullPath, out, 0666)
	}
	return err
}

func SaveStructureToJsonFile(folder string, filePath string, data interface{}) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var fullPath = fmt.Sprintf("%s%c%s.json", folder, os.PathSeparator, filePath)
	var out []byte
	out, err = json.Marshal(data)
	if err == nil {
		err = ioutil.WriteFile(fullPath, out, 0666)
	}
	return err
}

func LoadStructureFromYamlFile(folder string, filePath string, data interface{}) (interface{}, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var fullPath = fmt.Sprintf("%s%c%s.yaml", folder, os.PathSeparator, filePath)
	var out []byte
	var f *os.File
	f, err = os.Open(fullPath)
	if err == nil {
		defer f.Close()
		out, err = ioutil.ReadAll(f)
		if err == nil {
			err = yaml.Unmarshal(out, data)
		}
	}
	return data, err
}

func LoadStructureFromJsonFile(fullPath string, data interface{}) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	var out []byte
	var f *os.File
	f, err = os.Open(fullPath)
	if err == nil {
		defer f.Close()
		out, err = ioutil.ReadAll(f)
		if err == nil {
			err = json.Unmarshal(out, data)
		}
	}
	return err
}

func JsonToStructure(data string, itf interface{}) error {
	return json.Unmarshal([]byte(data), itf)
}

func StructureToJson(itf interface{}) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(buff)
	err := enc.Encode(itf)
	if err != nil {
		return []byte{}, err
	}
	return buff.Bytes(), nil

}

type Printable interface {
	String() string
}

func StructToString(i interface{}) string {
	if i == nil {
		return "<nil>"
	}
	switch v := i.(type) {
	case Printable:
		return v.String()
	default:
		return fmt.Sprintf("%v", i)
	}
}
