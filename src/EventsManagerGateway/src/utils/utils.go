package utils

import (
	"bytes"
	"encoding/gob"

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
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(obj)

	return b.Bytes()
}
