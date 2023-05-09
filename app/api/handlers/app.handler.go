package handlers

import "geoloc/app/services"

type Handler struct {
	UserHandler    *UserHandler
	AuthHandler    *AuthHandler
	GraphQLHandler *GraphQLHandler
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		NewUserHandler(services),
		NewAuthHandler(services),
		NewGraphQLHandler(services),
	}
}
