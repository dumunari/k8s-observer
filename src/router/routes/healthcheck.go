package routes

import (
	"net/http"
	"observer/src/controllers"
)

var healthCheckRoute = []Route{
	{
		Uri:      "/healthcheck",
		Method:   http.MethodGet,
		Function: controllers.HealthcheckController,
	},
}