package routes

import (
	"net/http"

	"github.com/dumunari/k8s-observer/src/api/controllers"
	"github.com/dumunari/k8s-observer/src/api/providers"
	"github.com/gorilla/mux"
)

var (
	deploymentsController controllers.DeploymentsController
	healthController      controllers.HealtcheckController
	nodesController       controllers.NodesController
)

func init() {

	responseUtils := providers.ProvideResponseUtils()

	nodesController = controllers.NodesController{
		NodesService:  providers.ProvideNodesService(),
		ResponseUtils: responseUtils,
	}

	healthController = controllers.HealtcheckController{
		ResponseUtils: responseUtils,
	}

	deploymentsController = controllers.DeploymentsController{
		DeploymentsService: providers.ProvideDeploymentsService(),
		ResponseUtils:      responseUtils,
	}
}

type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

func Configure(r *mux.Router) *mux.Router {
	routes := deploymentsRoutes
	routes = append(routes, healthCheckRoute[0])
	routes = append(routes, nodesRoutes[0])

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
