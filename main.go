package main

import (
	"github.com/felipecontra3/go-api-ping-pong/routes"
	"github.com/felipecontra3/go-api-ping-pong/servers"
)

func main() {
	r := routes.Configure()
	s := servers.CreateServer(r)
	s.ListenAndServe()
}
