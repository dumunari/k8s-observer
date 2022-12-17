package repositories

import (
	"fmt"
	"log"

	"github.com/dumunari/k8s-observer/src/infrastructure/client"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

type NodesRepository struct {
	Resources client.ResourcesInterface
}

type NodesRepositoryInterface interface {
	RetrieveNodes() (*v1.NodeList, error)
}

func (nodesRepositoryReceiver *NodesRepository) RetrieveNodes() (*v1.NodeList, error) {
	log.Println("[NodesRepository] - RetrieveNodes")

	nodesList, err := nodesRepositoryReceiver.Resources.RetrieveNodesList()

	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting nodes %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		return nil, err
	}

	return nodesList, nil
}
