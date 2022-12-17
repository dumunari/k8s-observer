package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dumunari/k8s-observer/src/api/utils"
	"github.com/dumunari/k8s-observer/src/models"
)

func TestJSONSuccess(t *testing.T) {
	//arrange
	writer := httptest.NewRecorder()
	statusCode := http.StatusOK
	data := models.Message{Message: "Test"}

	response := utils.Response{}

	//act
	response.JSON(writer, statusCode, data)
}

func TestError(t *testing.T) {
	//arrange
	writer := httptest.NewRecorder()
	statusCode := http.StatusInternalServerError
	errorInterface := errors.New("error")

	response := utils.Response{}

	//act
	response.Error(writer, statusCode, errorInterface)
}
