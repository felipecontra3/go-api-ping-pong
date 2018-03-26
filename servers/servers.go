package servers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// CreateServer returns a http.Server reference
func CreateServer(r *mux.Router) *http.Server {
	return &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}
}
