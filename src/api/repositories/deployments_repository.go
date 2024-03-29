package repositories

import (
	"fmt"
	"log"

	"github.com/dumunari/k8s-observer/src/infrastructure/client"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

type DeploymentsRepository struct {
	Resources client.ResourcesInterface
}

type DeploymentsRepositoryInterface interface {
	RetrieveDeployments() (*v1.DeploymentList, error)
}

func (deploymentsRepositoryReceiver *DeploymentsRepository) RetrieveDeployments() (*v1.DeploymentList, error) {
	log.Println("[DeploymentsRepository] - RetrieveDeployments")

	deploymentsList, err := deploymentsRepositoryReceiver.Resources.RetrieveDeploymentList()
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting deployments %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		return nil, err
	}

	return deploymentsList, nil
}
