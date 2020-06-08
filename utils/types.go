// Copyright 2020 Re-Bind Author (Fabrizio Torelli). All rights reserved.
// Use of this source code is governed by a LGPL-style
// license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// Error for Not Supported action/type/encoding
	ErrTypeNotSupport = errors.New("Type not supported")
	// Error for Not Invalid action/type/encoding
	ErrIPInvalid = errors.New("Invalid IP address")
)

// Convert a string to int typed value, and the success conversion flag
func StringToInt(v string) (int, bool) {
	if n, err := strconv.Atoi(v); err == nil {
		return n, true
	}
	return 0, false
}

// Convert a string to uint64 typed value, and the success conversion flag
func StringToUInt(v string, base int, bitSize int) (uint64, bool) {
	if n, err := strconv.ParseUint(v, base, bitSize); err == nil {
		return n, true
	}
	return 0, false
}

// Convert a string to long (int64) typed value, and the success conversion flag
func StringToLong(v string, base int, bitSize int) (int64, bool) {
	if n, err := strconv.ParseInt(v, base, bitSize); err == nil {
		return n, true
	}
	return 0, false
}

// Convert a string to float (float64) typed value, and the success conversion flag
func StringToFloat(v string, bitSize int) (float64, bool) {
	if n, err := strconv.ParseFloat(v, bitSize); err == nil {
		return n, true
	}
	return 0, false
}

// Convert a string to Datetime (time.Time), and reports eventually parsing errors
func StringToDateTime(v string, layout string) (time.Time, error) {
	return time.Parse(v, layout)
}

// Verify if a string item is present in a strings list
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

// Verify if an interface item is present in an interfaces list
func GenericListContainItem(elem interface{}, elems []interface{}) bool {
	for _, item := range elems {
		if item == elem {
			return true
		}
	}
	return false
}

// Replace not needed characters from a text string, to remove everithing is not a letter, digit, dot or dash
func ReplaceSimpleTextUnrelated(val string) (string, error) {
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

// Create a new Unique Identifier string
func NewUniqueIdentifier() string {
	return uuid.New().String()
}

func ConvertKeyToId(key string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(key, ".", "-"), " ", "-"))
}

func ConvertName(name string) string {
	return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(name), " ", "-"))
}
