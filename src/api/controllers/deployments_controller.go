package controllers

import (
	"log"
	"net/http"
	"observer/src/api/services"
	"observer/src/api/utils"
)

type DeploymentsController struct {
	DeploymentsService services.DeploymentsServiceInterface
	ResponseUtils      utils.ResponseInterface
}

func (modulesControllerReceiver *DeploymentsController) GetDeployments(writer http.ResponseWriter, _ *http.Request) {
	log.Println("[DeploymentsController] - GetDeployments")

	clusterDeployments, err := modulesControllerReceiver.DeploymentsService.RetrieveDeployments()
	if err != nil {
		modulesControllerReceiver.ResponseUtils.Error(writer, http.StatusInternalServerError, err)
		return
	}
	modulesControllerReceiver.ResponseUtils.JSON(writer, http.StatusOK, clusterDeployments)
}
