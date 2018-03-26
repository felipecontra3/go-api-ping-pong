package routes

import (
	"github.com/felipecontra3/go-api-ping-pong/ping"
	"github.com/gorilla/mux"
)

// Configure configura as rotas e retorna um Router
func Configure() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/ping", ping.Ping())
	return r
}
