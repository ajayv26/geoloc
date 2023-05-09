package services

import (
	"context"
	"database/sql"
	"fmt"
	"geoloc/app/helpers"
	"geoloc/app/models"
	"geoloc/app/stores"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofrs/uuid"
)

type AuthService struct {
	store *stores.Store
}

func NewAuthService(stores *stores.Store) *AuthService {
	return &AuthService{stores}
}

func (a *AuthService) GenerateJWT(email string) (string, error) {
	var mySecretKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *AuthService) GenerateOTP(ctx context.Context, tx *sql.Tx, userName string) (*models.OTPSession, error) {

	user, err := a.store.UserStore.GetByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	otpArg := models.OTPSession{
		UserID:    user.ID,
		Token:     helpers.GenerateRandomString(5),
		IsValid:   true,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(5)), // token will expire after 5 minutes
	}

	otp, err := stores.InsertOTPSessionToken(ctx, tx, &otpArg)
	if err != nil {
		return nil, err
	}
	return otp, nil

}

func (a *AuthService) Login(ctx context.Context, tx *sql.Tx, userName, otpToken string) (*models.AuthSession, error) {

	user, err := a.store.UserStore.GetByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	otp, err := stores.GetBySessionToken(ctx, otpToken)
	if err != nil {
		return nil, err
	}
	if !otp.IsValid {
		return nil, fmt.Errorf("otp is invalid")
	}

	if otp.ExpiresAt.UTC().Unix() < time.Now().UTC().Unix() {
		return nil, fmt.Errorf("session is expired")
	}
	if otp.UserID != user.ID {
		return nil, fmt.Errorf("otp is invalid")
	}

	otp.IsValid = false
	if err := stores.UpdateSessionToken(ctx, tx, otp); err != nil {
		return nil, err
	}

	return a.CreateAuthSession(ctx, tx, user.ID)

}

func (a *AuthService) CreateAuthSession(ctx context.Context, tx *sql.Tx, userID int64) (*models.AuthSession, error) {
	token, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	arg := &models.AuthSession{
		UserID:    userID,
		Token:     token,
		IsValid:   true,
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(720)), // token will expire after 5 minutes
	}

	return a.store.AuthStore.InsertAuthToken(ctx, tx, arg)
}
