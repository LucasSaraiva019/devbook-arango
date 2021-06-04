package routes

import (
	"net/http"

	"github.com/LucasSaraiva019/devbook-arango/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota represents all API routes
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

// Configure put all routes into the router
func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	return r
}
