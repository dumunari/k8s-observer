package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"observer/src/services"
	"observer/src/utils"
)

func GetServicesController(writer http.ResponseWriter, r *http.Request) {
	log.Println("GetServicesController")

	clusterServices, err := services.RetrieveServicesService()
	if err != nil {
		utils.Error(writer, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(writer, http.StatusOK, clusterServices)
}

func GetServicesByApplicationGroupController(writer http.ResponseWriter, r *http.Request) {
	log.Println("GetServicesByApplicationGroupController")

	params := mux.Vars(r)
	applicationGroup := params["applicationGroup"]

	clusterServices, err := services.RetrieveServicesByApplicationGroupService(applicationGroup)
	if err != nil {
		utils.Error(writer, http.StatusInternalServerError, err)
		return
	}

	if len(clusterServices) == 0 {
		err = errors.New(fmt.Sprintf("No services found for application group: %s", applicationGroup))
		utils.Error(writer, http.StatusNotFound, err)
		return
	}

	utils.JSON(writer, http.StatusOK, clusterServices)
}
