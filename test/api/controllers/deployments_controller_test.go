package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"observer/src/api/controllers"
	"observer/src/models"
	"testing"
)

func TestGetDeploymentsSuccess(t *testing.T){
	//arrange
	clusterDeployments := []models.Deployment{
		{
			Name: "deployment-one",
			RunningReplicas: 1,
			DesiredReplicas: 1,
			UnavailableReplicas: 0,
		},
		{
			Name: "deployment-two",
			RunningReplicas: 1,
			DesiredReplicas: 2,
			UnavailableReplicas: 1,
		},
	}
	retrieveDeployments = func() ([]models.Deployment, error) {
		return clusterDeployments, nil
	}

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/deployments", nil)
	statusCode := http.StatusOK

	dpSvc := &deploymentsServiceMock{}
	dpSvc.On("RetrieveDeployments")

	resUtils := &responseUtilsMock{}
	resUtils.On("JSON", writer, statusCode, clusterDeployments)

	dpCtrl := controllers.DeploymentsController{
		DeploymentsService: dpSvc,
		ResponseUtils:      resUtils,
	}

	dpCtrl.GetDeployments(writer, request)

	dpSvc.AssertCalled(t, "RetrieveDeployments")
	dpSvc.AssertNumberOfCalls(t, "RetrieveDeployments", 1)
	resUtils.AssertCalled(t, "JSON", writer, statusCode, clusterDeployments)
	resUtils.AssertNumberOfCalls(t, "JSON", 1)
	resUtils.AssertNotCalled(t, "Error")
}

func TestGetDeploymentsError(t *testing.T){
	//arrange
	var clusterDeployments []models.Deployment

	retrieveDeployments = func() ([]models.Deployment, error) {
		return clusterDeployments, errors.New("service error")
	}

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/deployments", nil)
	statusCode := http.StatusInternalServerError

	dpSvc := &deploymentsServiceMock{}
	dpSvc.On("RetrieveDeployments")

	resUtils := &responseUtilsMock{}
	resUtils.On("Error", writer, statusCode, errors.New("service error"))

	dpCtrl := controllers.DeploymentsController{
		DeploymentsService: dpSvc,
		ResponseUtils:      resUtils,
	}

	dpCtrl.GetDeployments(writer, request)

	dpSvc.AssertCalled(t, "RetrieveDeployments")
	dpSvc.AssertNumberOfCalls(t, "RetrieveDeployments", 1)
	resUtils.AssertCalled(t, "Error", writer, statusCode, errors.New("service error"))
	resUtils.AssertNumberOfCalls(t, "Error", 1)
	resUtils.AssertNotCalled(t, "JSON")
}