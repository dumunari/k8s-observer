package routes

import (
	"net/http"
)

var deploymentsRoutes = []Route{
	{
		Uri:      "/deployments",
		Method:   http.MethodGet,
		Function: deploymentsController.GetDeployments,
	},
}