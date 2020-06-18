package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
)

// Loada a given Rest Service request to a given structure, // Save a given structure to a given Rest Service response, acconding to the request Content-Type Cookies entry
func RestParseRequest(w http.ResponseWriter, r *http.Request, res interface{}) error {
	val := r.Header.Get("Content-Type")
	if val == "" {
		val = r.Header.Get("content-type")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if val == "" {
		val = http.DetectContentType(body)
	}
	if val == "" || strings.Index(strings.ToLower(val), "json") > 0 {
		//json response
		err = json.Unmarshal(body, res)
	} else if strings.Index(strings.ToLower(val), "yaml") > 0 {
		//json response
		err = yaml.Unmarshal(body, res)
	} else if strings.Index(strings.ToLower(val), "xml") > 0 {
		//json response
		err = xml.Unmarshal(body, res)
	} else {
		return errors.New(fmt.Sprintf("Unknown media type : %s", val))
	}
	if err != nil {
		return err
	}
	return nil
}

// Save a given structure to a given Rest Service response, acconding to the response Accept Cookies entry
func RestParseResponse(w http.ResponseWriter, r *http.Request, req interface{}) error {
	val := r.Header.Get("Accept")
	if val == "" {
		val = r.Header.Get("accept")
	}
	prettify := r.Header.Get("Prettify")
	if prettify == "" {
		prettify = r.Header.Get("prettify")
	}
	if prettify == "" {
		prettify = "false"
	}
	var byteArr []byte
	var err error
	if val == "" || strings.Index(strings.ToLower(val), "json") > 0 {
		//json response
		byteArr, err = json.Marshal(req)
		if prettify == "true" {
			b := bytes.Buffer{}
			json.Indent(&b, byteArr, "", "  ")
			byteArr = b.Bytes()
		}

		w.Header().Set("Content-Type", "application/json")
	} else if strings.Index(strings.ToLower(val), "yaml") > 0 {
		//json response
		byteArr, err = yaml.Marshal(req)
		w.Header().Set("Content-Type", "text/yaml")
	} else if strings.Index(strings.ToLower(val), "xml") > 0 {
		//json response
		if prettify == "true" {
			byteArr, err = xml.MarshalIndent(req, "", "  ")
		} else {
			byteArr, err = xml.Marshal(req)
		}
		w.Header().Set("Content-Type", "application/xml")
	} else {
		//json response
		byteArr, err = json.Marshal(req)
		if prettify == "true" {
			b := bytes.Buffer{}
			json.Indent(&b, byteArr, "", "  ")
			byteArr = b.Bytes()
		}

		w.Header().Set("Content-Type", "application/json")
		//		return errors.New(fmt.Sprintf("Unknown media type : %s", val))
	}
	if err != nil {
		return err
	}
	_, err = w.Write(byteArr)
	if err != nil {
		return err
	}
	return nil
}
