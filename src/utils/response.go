package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if jsonEncodeError := json.NewEncoder(writer).Encode(data); jsonEncodeError != nil {
		log.Fatal(jsonEncodeError)
	}
}

func Error(writer http.ResponseWriter, statusCode int, error error) {
	JSON(writer, statusCode, struct {
		Error string `json:"message"`
	}{
		Error: error.Error(),
	})
}
