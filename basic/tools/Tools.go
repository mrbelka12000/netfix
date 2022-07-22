package tools

import uuid "github.com/satori/go.uuid"

func GetRandomString() string {
	UUID := uuid.NewV4()
	return UUID.String()
}
