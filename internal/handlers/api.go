package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Rault647/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware (Middleware applied to all api endpoints)
	r.Use(chimiddle.StripSlashes) // Ignores trailing slashes in URL

	r.Route("/account", func(router chi.Router) {

		// Middleware for every request trying to access the /account route
		router.Use(middleware.Authorization) // Authorization function defined in internal/middleware directory

		router.Get("/coins", GetCoinBalance)	// GetCoinBalance function defined in internal/middleware directory
	})
}
