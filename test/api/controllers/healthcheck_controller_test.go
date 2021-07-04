package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"observer/src/api/controllers"
	"observer/src/models"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestGetHealthcheck(t *testing.T){
	//arrange
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/healthcheck", nil)
	message := models.Message{Message: "Server is up"}
	statusCode := http.StatusOK

	responseUtils := &responseUtilsMock{}
	responseUtils.On("JSON", writer, statusCode, message)

	healthcheckController := &controllers.HealtcheckController{
		ResponseUtils: responseUtils,
	}

	//act
	healthcheckController.GetHealthcheck(writer, request)

	//assert
	responseUtils.AssertCalled(t, "JSON", writer, statusCode, message)
}


