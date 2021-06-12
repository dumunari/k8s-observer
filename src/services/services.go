package services

import (
	"log"
	"observer/src/models"
	"observer/src/repository"
)

func RetrieveServicesService() ([]models.Service, error) {
	log.Println("RetrieveServicesService")

	var services []models.Service
	deployments, err := repository.RetrieveDeploymentsRepository()
	if err != nil {
		return services, err
	}

	for _, deployment := range deployments.Items {
		service := models.Service{
			Name: deployment.Spec.Template.Labels["service"],
			ApplicationGroup: deployment.Labels["applicationGroup"],
			RunningPodsCount: deployment.Status.AvailableReplicas,
		}
		services = append(services, service)
	}
	return services, nil
}

func RetrieveServicesByApplicationGroupService(applicationGroup string) ([]models.Service, error) {
	log.Println("RetrieveServicesByApplicationGroupService")

	var services []models.Service
	deployments, err := repository.RetrieveDeploymentsRepository()
	if err != nil {
		return services, err
	}

	for _, deployment := range deployments.Items {
		if deployment.Labels["applicationGroup"] != "" && deployment.Labels["applicationGroup"] == applicationGroup {
			service := models.Service{
				Name: deployment.Name,
				ApplicationGroup: deployment.Labels["applicationGroup"],
				RunningPodsCount: deployment.Status.AvailableReplicas,
			}
			services = append(services, service)
		} else if deployment.Labels["applicationGroup"] == "" && applicationGroup == "none" {
			service := models.Service{
				Name: deployment.Name,
				ApplicationGroup: deployment.Labels["applicationGroup"],
				RunningPodsCount: deployment.Status.AvailableReplicas,
			}
			services = append(services, service)
		}
	}
	return services, nil
}