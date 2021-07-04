package router

import (
	"github.com/gorilla/mux"
	"observer/src/api/router/routes"
)

func GenerateRouter() *mux.Router {
	return routes.Configure(mux.NewRouter())
}
