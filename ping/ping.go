package ping

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Body string `json:"body"`
}

// Ping responde a requisição em /ping com json
// com body igual a "pong"
func Ping() http.Handler {
	return http.HandlerFunc(pingHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{"pong"})
}
