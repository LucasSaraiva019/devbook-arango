package middlewares

import (
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}
