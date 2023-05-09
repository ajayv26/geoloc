package routes

import (
	"geoloc/app/api/handlers"

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
		rt.GraphQl(r)
	})
}

func (rt *Route) UserRoutes(r chi.Router) {
	h := rt.handler.UserHandler
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.UserListHandler)
		r.Get("/{id}", h.UserGetByIDHandler)
		r.Post("/", h.UserInsertHandler)
		r.Put("/{id}", h.UserUpdateHandler)
		r.Delete("/{id}", h.DeleteUserHandler)
	})
}

func (rt *Route) AuthRoutes(r chi.Router) {
	h := rt.handler.AuthHandler
	r.Route("/auth", func(r chi.Router) {
		r.Post("/otp", h.GetOTPHandler)
		r.Post("/login", h.LoginHandler)
	})
}

func (rt *Route) GraphQl(r chi.Router) {
	h := rt.handler.GraphQLHandler
	r.Route("/gql", func(r chi.Router) {
		r.Handle("/", h.PlayGround())
		r.Handle("/query", h.Query())
	})
}
