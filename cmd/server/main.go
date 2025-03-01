package main

import (
	"github.com/y3g0r/rehearsals-api-go/internal/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	cfg := config.LoadConfig()

	http.ListenAndServe(cfg.ServerAddress, r)
}
