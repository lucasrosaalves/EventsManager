package utils

import (
	"encoding/json"
)

func GenerateBytes(obj interface{}) []byte {
	payload, _ := json.Marshal(obj)

	return payload
}
