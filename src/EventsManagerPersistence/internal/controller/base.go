package controller

import (
	"encoding/json"
	"net/http"
)

func Ok(writer http.ResponseWriter, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	if object == nil {
		return
	}

	content, _ := json.Marshal(object)
	writer.Write(content)
}
