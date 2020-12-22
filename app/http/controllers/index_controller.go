package controllers

import "net/http"

type IndexController struct {
}

func (*IndexController) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
	w.Write([]byte("Index:index"))
}
