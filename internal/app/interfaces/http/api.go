package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API struct
type API struct {
	router *chi.Mux
}

// NewAPI create a new api
func NewAPI() *API {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello world"))
	})

	return &API{router}
}

// ListenAndServe ..
func (api *API) ListenAndServe(hostAndPort string) error {
	if strings.HasPrefix(hostAndPort, ":") {
		hostAndPort = "localhost" + hostAndPort
	}

	log.Printf("Run server on %v\n", hostAndPort)
	return http.ListenAndServe(hostAndPort, api.router)
}
