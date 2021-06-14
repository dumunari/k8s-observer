package repository

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"observer/src/repository/infrastructure"
)


type ClientInterface interface {
	RetrieveDeploymentList() (*v1.DeploymentList, error)
}

type DeploymentsRepository struct{
	Clientset kubernetes.Interface
}

type DeploymentsRepositoryInterface interface {
	RetrieveDeployments() (*v1.DeploymentList, error)
}

func ProvideDeploymentsRepository() *DeploymentsRepository {
	return &DeploymentsRepository{
		Clientset: infrastructure.ProvideClusterConfig().RetrieveClientSet(),
	}
}

func (deploymentsRepositoryInterface *DeploymentsRepository) RetrieveDeployments() (*v1.DeploymentList, error) {
	log.Println("RetrieveDeploymentsRepository")

	deploymentsList, err := deploymentsRepositoryInterface.Clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting deployments %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		return nil, err
	}

	return deploymentsList, nil
}
