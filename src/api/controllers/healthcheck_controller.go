package controllers

import (
	"log"
	"net/http"

	"github.com/dumunari/k8s-observer/src/api/utils"
	"github.com/dumunari/k8s-observer/src/models"
)

type HealtcheckController struct {
	ResponseUtils utils.ResponseInterface
}

func (healthcheckControllerReceiver *HealtcheckController) GetHealthcheck(writer http.ResponseWriter, _ *http.Request) {
	log.Println("[HealtcheckController] - GetHealthcheck")

	healthcheckControllerReceiver.ResponseUtils.JSON(writer, http.StatusOK, models.Message{
		Message: "Server is up",
	})
}
