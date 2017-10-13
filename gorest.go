package gorest

import (
	"bytes"
	"fmt"
)

/*
MapToQueryString turns a map of [string]string to a string of query parameters
Example MapToQueryString({"a":"b"}) -> ?a=b
*/
func MapToQueryString(m map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for k, v := range m {
		buffer.WriteString(fmt.Sprintf("%s=%s", k, v))
	}
	return buffer.String()
}
