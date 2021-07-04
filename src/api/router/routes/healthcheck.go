package routes

import (
	"net/http"
)

var healthCheckRoute = []Route{
	{
		Uri:      "/healthcheck",
		Method:   http.MethodGet,
		Function: healthController.GetHealthcheck,
	},
}