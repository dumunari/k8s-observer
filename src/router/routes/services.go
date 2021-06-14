package routes

import (
	"net/http"
	"observer/src/controllers"
)

var servicesRoutes = []Route{
	{
		Uri:      "/services",
		Method:   http.MethodGet,
		Function: controllers.IServicesController.GetServicesController,
	},
	{
		Uri:      "/services/{applicationGroup}",
		Method:   http.MethodGet,
		Function: controllers.IServicesController.GetServicesByApplicationGroupController,
	},
}