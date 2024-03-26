package handlers

import (
	"github.com/StanleyAnyas/goapi/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(route chi.Router) {
		route.Use(middleware.Authorization)

		route.Get("/coins", GetCoinBalance)
	})
}
