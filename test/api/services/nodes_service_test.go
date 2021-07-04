package services

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"observer/src/api/services"
	"testing"
)

func TestRetrieveNodesSuccess(t *testing.T) {
	//arrange
	retrieveNodes = func() (*v1.NodeList, error) {
		return &v1.NodeList{
			Items: []v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node1",
					},
					Status: v1.NodeStatus{
						Conditions: []v1.NodeCondition{
							{
								Status: "False",
							},
							{
								Status: "False",
							},
							{
								Status: "False",
							},
							{
								Status: "True",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node2",
					},
					Status: v1.NodeStatus{
						Conditions: []v1.NodeCondition{
							{
								Status: "True",
							},
							{
								Status: "False",
							},
							{
								Status: "False",
							},
							{
								Status: "False",
							},
						},
					},
				},
			},
		}, nil
	}

	repository := &nodeRepositoryMock{}
	repository.On("RetrieveNodes")

	nodesService := &services.NodesService{
		Repository: repository,
	}

	//act
	clusterNodes, err := nodesService.RetrieveNodes()

	//assert
	assert.EqualValues(t, 2, len(clusterNodes))
	assert.Nil(t, err)
}

func TestRetrieveNodesError(t *testing.T) {
	//arrange
	retrieveNodes = func() (*v1.NodeList, error) {
		return &v1.NodeList{}, errors.New("repository error")
	}

	repository := &nodeRepositoryMock{}
	repository.On("RetrieveNodes")

	nodesService := &services.NodesService{
		Repository: repository,
	}

	//act
	clusterNodes, err := nodesService.RetrieveNodes()

	//assert
	assert.EqualValues(t, "repository error", fmt.Sprint(err))
	assert.Nil(t, clusterNodes)
}
