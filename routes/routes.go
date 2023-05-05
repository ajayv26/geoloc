package routes

import (
	"geoloc/handlers"

	"github.com/go-chi/chi"
)

type Route struct {
	handler *handlers.Handler
}

func NewRoute(handler *handlers.Handler) *Route {
	return &Route{handler}
}

func Routes(router *chi.Mux, rt *Route) {
	// group other routes with /api
	router.Route("/api", func(r chi.Router) {
		rt.UserRoutes(r)
		rt.AuthRoutes(r)
	})
}

func (rt *Route) UserRoutes(r chi.Router) {
	h := handlers.UserHandler{}
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.UserListHandler)
		r.Get("/{id}", h.UserGetByIDHandler)
		r.Post("/", h.UserInsertHandler)
		r.Put("/{id}", h.UserUpdateHandler)
		r.Delete("/{id}", h.DeleteUserHandler)
	})
}

func (rt *Route) AuthRoutes(r chi.Router) {
	h := handlers.AuthHandler{}
	r.Route("/auth", func(r chi.Router) {
		r.Post("/otp", h.GetOTPHandler)
		r.Post("/login", h.LoginHandler)
	})
}
