package controllers

import (
	"log"
	"net/http"
	"observer/src/api/utils"
	"observer/src/models"
)

type HealtcheckController struct {
	ResponseUtils      utils.ResponseInterface
}

func (healthcheckControllerReceiver *HealtcheckController) GetHealthcheck(writer http.ResponseWriter, _ *http.Request) {
	log.Println("[HealtcheckController] - GetHealthcheck")

	healthcheckControllerReceiver.ResponseUtils.JSON(writer, http.StatusOK, models.Message{
		Message: "Server is up",
	})
}