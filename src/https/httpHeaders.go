/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package https

import (
    "strings"
)


func SetHeaderIfNotSet(headers map[string]string, name, value string) {
	if headers == nil {
		headers = make(map[string]string)
		headers[name] = value
	} else if headers[name] == "" {
		headers[name] = value
	}
}

func SetHeaders(headers map[string]string, name, value string) {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers[name] = value
}

func MergeHeaders(headers, headersToMerge map[string]string) {
	for k, v := range headersToMerge {
		if _, ok := headers[k]; !ok {
			headers[k] = v
		}
	}
}

func LookupCaseInsensitive(obj map[string]string, prop string) *string {
	prop = strings.ToLower(prop)
	for k := range obj {
		if prop == strings.ToLower(k) {
			return &k
		}
	}
	return nil
}
