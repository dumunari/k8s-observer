package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct{
}

type ResponseInterface interface {
	JSON(http.ResponseWriter, int, interface{})
	Error(http.ResponseWriter, int, error)
}

func (responseInterface *Response) JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if jsonEncodeError := json.NewEncoder(writer).Encode(data); jsonEncodeError != nil {
		log.Fatal(jsonEncodeError)
	}
}

func (responseInterface *Response) Error(writer http.ResponseWriter, statusCode int, error error) {
	responseInterface.JSON(writer, statusCode, struct {
		Error string `json:"message"`
	}{
		Error: error.Error(),
	})
}
