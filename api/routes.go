package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router = chi.NewRouter()

// Router returns the api router
func Router() http.Handler {
	router.Use(middleware.Logger)
	router.Use(recoverer)
	router.Use(capture)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			Status: http.StatusOK,
			Data:   "Congratulations - Service running...",
		}
		resp.ServerJSON(w)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			Status: http.StatusOK,
			Data:   "route not found",
		}
		resp.ServerJSON(w)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			Status: http.StatusOK,
			Data:   "method not allowed",
		}
		resp.ServerJSON(w)
	})

	registerRoutes()

	return router
}

func registerRoutes() {
	router.Route("/v1", func(r chi.Router) {
		r.Get("/", index)

		r.Mount("/monkeys", monkeyRoutes())
	})
}

func monkeyRoutes() http.Handler {
	hr := NewMonkeyRoutes()
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Post("/", hr.saveMonkey)
		r.Get("/", hr.listMonkey)
		r.Get("/{id}", hr.getMonkey)
		r.Put("/{id}", hr.updateMonkey)
		r.Delete("/{id}", hr.deleteMonkey)
	})
	return h
}
