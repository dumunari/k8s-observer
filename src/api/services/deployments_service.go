package services

import (
	"log"

	"github.com/dumunari/k8s-observer/src/api/repositories"
	"github.com/dumunari/k8s-observer/src/models"
)

type DeploymentsService struct {
	Repository repositories.DeploymentsRepositoryInterface
}

type DeploymentsServiceInterface interface {
	RetrieveDeployments() ([]models.Deployment, error)
}

func (deploymentsServiceReceiver *DeploymentsService) RetrieveDeployments() ([]models.Deployment, error) {
	log.Println("[DeploymentsService] - RetrieveDeployments")

	var deployments []models.Deployment
	clusterDeployments, err := deploymentsServiceReceiver.Repository.RetrieveDeployments()
	if err != nil {
		return deployments, err
	}

	for _, clusterDeployment := range clusterDeployments.Items {
		deployment := models.Deployment{
			Name:                clusterDeployment.Name,
			RunningReplicas:     clusterDeployment.Status.AvailableReplicas,
			UnavailableReplicas: clusterDeployment.Status.UnavailableReplicas,
			DesiredReplicas:     clusterDeployment.Status.Replicas,
		}
		deployments = append(deployments, deployment)
	}
	return deployments, nil
}
