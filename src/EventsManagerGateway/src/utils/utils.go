package utils

import (
	"encoding/json"

	uuid "github.com/nu7hatch/gouuid"
)

func GenerateGuid() string {
	u, err := uuid.NewV4()

	if err != nil {
		return ""
	}

	return u.String()
}

func GenerateBytes(obj interface{}) []byte {
	payload, _ := json.Marshal(obj)

	return payload
}
