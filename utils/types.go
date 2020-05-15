// Copyright 2020 Re-Bind Author (Fabrizio Torelli). All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrTypeNotSupport = errors.New("Type not supported")
	ErrIPInvalid      = errors.New("Invalid IP address")
)

func getInteger(v string) (int, bool) {
	if n, err := strconv.Atoi(v); err == nil {
		return n, true
	}
	return 0, false
}

func StringsListContainItem(elem string, elems []string, ignoreCase bool) bool {
	if ignoreCase {
		elemLower := strings.ToLower(elem)
		for _, item := range elems {
			if ignoreCase && strings.ToLower(item) == elemLower {
				return true
			}
		}
	} else {
		for _, item := range elems {
			if item == elem {
				return true
			}
		}
	}
	return false
}

func GenericListContainItem(elem interface{}, elems []interface{}) bool {
	for _, item := range elems {
		if item == elem {
			return true
		}
	}
	return false
}

func ReplaceQuestionUnrelated(val string) (string, error) {
	expr, err := regexp.Compile("[^a-zA-Z0-9.-]+")
	if err != nil {
		return val, err
	}
	value := expr.ReplaceAllString(val, "")
	if value[len(value)-1:] == "." {
		value = value[:len(value)-1]
	}
	return value, nil
}

func ConvertKeyToId(key string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(key, ".", "-"), " ", "-"))
}
