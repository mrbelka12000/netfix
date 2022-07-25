package tools

import (
	"bytes"
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetRandomString() string {
	UUID := uuid.NewV4()
	return UUID.String()
}

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

func GetUnixDate() int64 {
	loc, _ := time.LoadLocation("Asia/Almaty")
	now := time.Now().In(loc).Unix()

	return now
}
