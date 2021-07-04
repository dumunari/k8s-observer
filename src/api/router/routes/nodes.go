package routes

import (
	"net/http"
)

var nodesRoutes = []Route{
	{
		Uri:      "/nodes",
		Method:   http.MethodGet,
		Function: nodesController.GetNodes,
	},
}