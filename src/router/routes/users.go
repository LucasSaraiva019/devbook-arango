package routes

import (
	"net/http"

	"github.com/LucasSaraiva019/devbook-arango/src/controllers"
)

var routesUsers = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
	},
}
