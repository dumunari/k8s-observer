package router

import (
	"github.com/dumunari/k8s-observer/src/api/router/routes"
	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	return routes.Configure(mux.NewRouter())
}
