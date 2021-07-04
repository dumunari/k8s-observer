package services

import (
	"github.com/stretchr/testify/mock"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

var (
	retrieveDeployments func() (*v12.DeploymentList, error)
	retrieveNodes func() (*v1.NodeList, error)
)

type deploymentsRepositoryMock struct {
	mock.Mock
}

func (deploymentsRepositoryMock *deploymentsRepositoryMock) RetrieveDeployments() (*v12.DeploymentList, error) {
	deploymentsRepositoryMock.Called()
	return retrieveDeployments()
}

type nodeRepositoryMock struct {
	mock.Mock
}

func (nodeRepositoryMock *nodeRepositoryMock) RetrieveNodes() (*v1.NodeList, error){
	nodeRepositoryMock.Called()
	return retrieveNodes()
}
