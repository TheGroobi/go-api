package handlers

import (
	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	//Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {

		//Middleware for /account route
		router.use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
