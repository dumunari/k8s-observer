package controllers

import (
	"log"
	"net/http"
	"observer/src/api/services"
	"observer/src/api/utils"
)

type NodesController struct{
	NodesService  services.NodesServiceInterface
	ResponseUtils utils.ResponseInterface
}

func (nodesControllerReceiver *NodesController) GetNodes(writer http.ResponseWriter, _ *http.Request) {
	log.Println("[NodesController] - GetNodes")

	clusterNodes, err := nodesControllerReceiver.NodesService.RetrieveNodes()
	if err != nil {
		nodesControllerReceiver.ResponseUtils.Error(writer, http.StatusInternalServerError, err)
		return
	}
	nodesControllerReceiver.ResponseUtils.JSON(writer, http.StatusOK, clusterNodes)
}
