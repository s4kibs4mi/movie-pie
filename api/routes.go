package api

import (
	"github.com/codersgarage/golang-restful-boilerplate/api/rest"
	"github.com/codersgarage/golang-restful-boilerplate/api/rpc"
	"github.com/codersgarage/golang-restful-boilerplate/proto/defs"
	"github.com/twitchtv/twirp"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router = chi.NewRouter()

// Router returns the api router
func Router() http.Handler {
	router.Use(middleware.Logger)
	router.Use(rest.Recoverer)
	router.Use(rest.Capture)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp := rest.Response{
			Status: http.StatusOK,
			Data:   "Congratulations - Service running...",
		}
		resp.ServerJSON(w)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		resp := rest.Response{
			Status: http.StatusOK,
			Data:   "route not found",
		}
		resp.ServerJSON(w)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		resp := rest.Response{
			Status: http.StatusOK,
			Data:   "method not allowed",
		}
		resp.ServerJSON(w)
	})

	registerRoutes()

	return router
}

func registerRoutes() {
	rpcServer := &rpc.RPCServer{}
	rpcHandler := defs.NewMonkeyRPCServiceServer(rpcServer, twirp.ChainHooks(rpc.AuthHook()))

	router.Mount(defs.MonkeyRPCServicePathPrefix, rpcHandler)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/", rest.Index)
		r.Mount("/monkeys", monkeyRoutes())
	})
}

func monkeyRoutes() http.Handler {
	hr := rest.NewMonkeyRoutes()
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Post("/", hr.SaveMonkey)
		r.Get("/", hr.ListMonkey)
		r.Get("/{id}", hr.GetMonkey)
		r.Put("/{id}", hr.UpdateMonkey)
		r.Delete("/{id}", hr.DeleteMonkey)
	})
	return h
}
