package gorest

import (
	"bytes"
	"fmt"
	"strings"
)

/*
MapToQueryString turns a map of [string]string to a string of query parameters
Example MapToQueryString({"a":"b"}) -> ?a=b
*/
func MapToQueryString(m map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for k, v := range m {
		buffer.WriteString(fmt.Sprintf("%s=%s&", k, v))
	}
	bufferString := buffer.String()
	return bufferString[:len(bufferString)-1]
}

/*
Join joins parts of a URL including the query string
*/
func Join(parts []string) string {
	url := strings.Trim(parts[0], "/")
	parts = parts[1:]
	for _, part := range parts {
		cleanPart := strings.Trim(part, "/")
		if strings.HasPrefix(cleanPart, "?") {
			url = strings.Join([]string{url, cleanPart}, "")
			continue
		}
		url = strings.Join([]string{url, cleanPart}, "/")
	}
	return url
}
