package services

import (
	"log"
	"observer/src/api/repositories"
	"observer/src/models"
)

type NodesService struct {
	Repository repositories.NodesRepositoryInterface
}

type NodesServiceInterface interface {
	RetrieveNodes() ([]models.Node, error)
}

func (nodesServiceInterface *NodesService) RetrieveNodes() ([]models.Node, error) {
	log.Println("[NodesService] - RetrieveNodes")

	var nodes []models.Node
	clusterNodes, err := nodesServiceInterface.Repository.RetrieveNodes()
	if err != nil {
		return nodes, err
	}

	for _, clusterNode := range clusterNodes.Items {
		node := models.Node{
			Name: clusterNode.Name,
			MemoryPressure: string(clusterNode.Status.Conditions[0].Status),
			DiskPressure: string(clusterNode.Status.Conditions[1].Status),
			PIDPressure: string(clusterNode.Status.Conditions[2].Status),
			Ready: string(clusterNode.Status.Conditions[3].Status),
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}