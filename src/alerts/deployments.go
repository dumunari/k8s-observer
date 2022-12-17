package alerts

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dumunari/k8s-observer/src/models"
	v1 "k8s.io/api/apps/v1"
)

func DeploymentsCheck() {
	log.Println("[Alerts] - DeploymentsCheck")

	deploymentsList, err := resources.RetrieveDeploymentList()
	if err != nil {
		log.Println("[Alerts] - Error on resources.RetrieveDeploymentList")
		return
	}

	for _, deploymentInList := range deploymentsList.Items {
		if hasUnavailableReplicas(deploymentInList) {
			deploymentInfo := &models.Deployment{
				Name:                deploymentInList.Name,
				RunningReplicas:     deploymentInList.Status.AvailableReplicas,
				UnavailableReplicas: deploymentInList.Status.UnavailableReplicas,
				DesiredReplicas:     deploymentInList.Status.Replicas,
			}
			info, _ := json.Marshal(deploymentInfo)
			fmt.Println(string(info))
		}
	}
}

func hasUnavailableReplicas(deployment v1.Deployment) bool {
	return deployment.Status.UnavailableReplicas != 0
}
