package routes

import (
	"github.com/gorilla/mux"
	"goframwork/app/http/controllers"
)

func MapWebRoutes(r *mux.Router) {

	a := new(controllers.IndexController)
	r.HandleFunc("/", a.Index)
}
