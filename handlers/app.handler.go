package handlers

import "geoloc/services"

type Handler struct {
	UserHandler *UserHandler
	AuthHandler *AuthHandler
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		NewUserHandler(services),
		NewAuthHandler(services),
	}
}
