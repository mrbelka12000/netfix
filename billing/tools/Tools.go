package tools

import (
	"bytes"
	"encoding/json"
)

func MakeJsonString(value interface{}) string {
	if value == nil {
		return "{}"
	}
	bf := bytes.NewBufferString("")
	e := json.NewEncoder(bf)
	e.SetEscapeHTML(false)
	e.Encode(value)
	return bf.String()
}

func PtrBool(v bool) *bool {
	return &v
}
