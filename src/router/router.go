package router

import (
	"github.com/gorilla/mux"
	"observer/src/router/routes"
)

func GenerateRouter() *mux.Router {
	return routes.Configure(mux.NewRouter())
}
