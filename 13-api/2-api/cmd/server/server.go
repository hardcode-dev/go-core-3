package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-dev-v3/13-api/2-api/pkg/api"
)

type server struct {
	api    *api.API
	router *mux.Router
}

func main() {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = api.New(srv.router)
	srv.api.Endpoints()
	http.ListenAndServe(":8080", srv.router)
}
