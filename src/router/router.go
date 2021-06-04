package router

import (
	"github.com/LucasSaraiva019/devbook-arango/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return a router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
