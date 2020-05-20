package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

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
		out, err = ioutil.ReadAll(f)
		if err == nil {
			err = yaml.Unmarshal(out, data)
		}
	}
	return data, err
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

type Printable interface{
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