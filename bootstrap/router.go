package bootstrap

import (
	"github.com/gorilla/mux"
	"goframwork/routes"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	routes.MapWebRoutes(router)

	return router
}
