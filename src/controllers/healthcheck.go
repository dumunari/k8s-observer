package controllers

import (
	"log"
	"net/http"
	"observer/src/models"
	"observer/src/utils"
)

func HealthcheckController(writer http.ResponseWriter, r *http.Request) {
	log.Println("HealthcheckController")

	utils.IResponse.JSON(writer, http.StatusOK, models.Message{
		Message: "Server is up",
	})
}