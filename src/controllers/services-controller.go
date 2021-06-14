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

var (
	IServicesController servicesControllerInterface = &ServicesController{}
)

type ServicesController struct{
	Service services.ServicesServiceInterface
	ResponseUtils utils.ResponseInterface
}

type servicesControllerInterface interface {
	GetServicesController(http.ResponseWriter, *http.Request)
	GetServicesByApplicationGroupController(http.ResponseWriter, *http.Request)
}

func (servicesControllerInterface *ServicesController) GetServicesController(writer http.ResponseWriter, r *http.Request) {
	log.Println("GetServicesController")

	servicesControllerInterface.Service = services.ProvideServicesService()
	servicesControllerInterface.ResponseUtils = utils.ProvideResponseUtils()

	clusterServices, err := servicesControllerInterface.Service.RetrieveServicesService()
	if err != nil {
		servicesControllerInterface.ResponseUtils.Error(writer, http.StatusInternalServerError, err)
		return
	}
	servicesControllerInterface.ResponseUtils.JSON(writer, http.StatusOK, clusterServices)
}

func (servicesControllerInterface *ServicesController) GetServicesByApplicationGroupController(writer http.ResponseWriter, r *http.Request) {
	log.Println("GetServicesByApplicationGroupController")

	params := mux.Vars(r)
	applicationGroup := params["applicationGroup"]

	servicesControllerInterface.Service = services.ProvideServicesService()
	servicesControllerInterface.ResponseUtils = utils.ProvideResponseUtils()

	clusterServices, err := servicesControllerInterface.Service.RetrieveServicesByApplicationGroupService(applicationGroup)
	if err != nil {
		servicesControllerInterface.ResponseUtils.Error(writer, http.StatusInternalServerError, err)
		return
	}

	if len(clusterServices) == 0 {
		err = errors.New(fmt.Sprintf("No services found for application group: %s", applicationGroup))
		servicesControllerInterface.ResponseUtils.Error(writer, http.StatusNotFound, err)
		return
	}
	servicesControllerInterface.ResponseUtils.JSON(writer, http.StatusOK, clusterServices)
}
