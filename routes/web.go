package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MapWebRoutes(r *mux.Router) {

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("首页"))
	})
}
