package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/vipin-jose/goapi/internal/middleware"
)

// If name of function starts with a capital letter, it is exported (can be imported from other packages)
// If it small latter, it is a private function
func Handler(r *chi.Mux) {
	r.Use(chimiddle.Logger)       // Log every request
	r.Use(chimiddle.StripSlashes) // Strip slashes from the request
	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization) // Use the Authorization middleware
		r.Get("/coins", GetCoinBalance) // Handle GET request to /api/coins
	})
}

//Middleware is a function that gets called before the request is passed to the handler
