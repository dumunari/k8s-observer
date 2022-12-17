package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dumunari/k8s-observer/src/api/controllers"
	"github.com/dumunari/k8s-observer/src/models"
)

func TestGetNodesSuccess(t *testing.T) {
	//arrange
	clusterNodes := []models.Node{
		{
			Name:           "deployment-one",
			DiskPressure:   "False",
			MemoryPressure: "False",
			PIDPressure:    "False",
			Ready:          "True",
		},
		{
			Name:           "deployment-one",
			DiskPressure:   "True",
			MemoryPressure: "False",
			PIDPressure:    "False",
			Ready:          "False",
		},
	}
	retrieveNodes = func() ([]models.Node, error) {
		return clusterNodes, nil
	}

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/nodes", nil)
	statusCode := http.StatusOK

	ndSvc := &nodesServiceMock{}
	ndSvc.On("RetrieveNodes")

	resUtils := &responseUtilsMock{}
	resUtils.On("JSON", writer, statusCode, clusterNodes)

	dpCtrl := controllers.NodesController{
		NodesService:  ndSvc,
		ResponseUtils: resUtils,
	}

	dpCtrl.GetNodes(writer, request)

	ndSvc.AssertCalled(t, "RetrieveNodes")
	ndSvc.AssertNumberOfCalls(t, "RetrieveNodes", 1)
	resUtils.AssertCalled(t, "JSON", writer, statusCode, clusterNodes)
	resUtils.AssertNumberOfCalls(t, "JSON", 1)
	resUtils.AssertNotCalled(t, "Error")
}

func TestGetNodesError(t *testing.T) {
	//arrange
	var clusterNodes []models.Node

	retrieveNodes = func() ([]models.Node, error) {
		return clusterNodes, errors.New("service error")
	}

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/nodes", nil)
	statusCode := http.StatusInternalServerError

	ndSvc := &nodesServiceMock{}
	ndSvc.On("RetrieveNodes")

	resUtils := &responseUtilsMock{}
	resUtils.On("Error", writer, statusCode, errors.New("service error"))

	ndCtrl := controllers.NodesController{
		NodesService:  ndSvc,
		ResponseUtils: resUtils,
	}

	ndCtrl.GetNodes(writer, request)

	ndSvc.AssertCalled(t, "RetrieveNodes")
	ndSvc.AssertNumberOfCalls(t, "RetrieveNodes", 1)
	resUtils.AssertCalled(t, "Error", writer, statusCode, errors.New("service error"))
	resUtils.AssertNumberOfCalls(t, "Error", 1)
	resUtils.AssertNotCalled(t, "JSON")
}
