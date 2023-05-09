package handlers

import (
	"encoding/json"
	"fmt"
	"geoloc/app/models"
	"geoloc/app/services"
	"geoloc/app/stores"
	"net/http"

	"github.com/go-chi/render"
)

type AuthHandler struct {
	service *services.Service
}

func NewAuthHandler(service *services.Service) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) GetOTPHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.OTPLoginRequest{}
	if decodErr := json.NewDecoder(r.Body).Decode(&req); decodErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	defer stores.RollbackTx(ctx, tx)

	obj, err := h.service.AuthService.GenerateOTP(ctx, tx, req.UserName)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, obj.Token)
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.OTPLoginRequest{}
	if decodErr := json.NewDecoder(r.Body).Decode(&req); decodErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	tx, err := stores.BeginTx(ctx)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	defer stores.RollbackTx(ctx, tx)

	obj, err := h.service.AuthService.Login(ctx, tx, req.UserName, req.OTP)
	if err != nil {
		render.JSON(w, r, err.Error())
		return
	}

	if err := stores.CommitTx(ctx, tx); err != nil {
		render.JSON(w, r, err.Error())
		return
	}
	render.JSON(w, r, obj)

}
