package controllers

import (
	"github.com/stretchr/testify/mock"
	"net/http"
	"observer/src/models"
)

var retrieveDeployments func() ([]models.Deployment, error)
var retrieveNodes func() ([]models.Node, error)

type deploymentsServiceMock struct{
	mock.Mock
}

func (deploymentsServiceMock *deploymentsServiceMock) RetrieveDeployments() ([]models.Deployment, error) {
	deploymentsServiceMock.Called()
	return retrieveDeployments()
}

type nodesServiceMock struct {
	mock.Mock
}

func (nodesServiceMock *nodesServiceMock) RetrieveNodes() ([]models.Node, error) {
	nodesServiceMock.Called()
	return retrieveNodes()
}

type responseUtilsMock struct{
	mock.Mock
}

func (mock *responseUtilsMock) JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	mock.Called(writer, statusCode, data)
}

func (mock *responseUtilsMock) Error(writer http.ResponseWriter, statusCode int, error error) {
	mock.Called(writer, statusCode, error)
}