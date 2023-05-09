package services

import (
	"geoloc/app/stores"
)

type Service struct {
	UserService *UserService
	AuthService *AuthService
}

func NewService(store *stores.Store) *Service {
	return &Service{
		NewUserService(store),
		NewAuthService(store),
	}
}
