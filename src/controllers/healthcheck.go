package controllers

import (
	"log"
	"net/http"
	"observer/src/models"
	"observer/src/utils"
)

func HealthcheckController(writer http.ResponseWriter, r *http.Request) {
	log.Println("HealthcheckController")

	healthcheck := models.Healthcheck{
		Message: "Server is up",
	}

	utils.JSON(writer, http.StatusOK, healthcheck)
}