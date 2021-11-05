package services

import (
	"github.com/orlovssky/gread/internal/secrets"
	"github.com/orlovssky/gread/internal/store"
	"github.com/orlovssky/gread/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

var authStore store.AuthStore

func (s *AuthService) SignIn(email, password string) (string, error) {
	// Check user with that email exists
	user, err := authStore.SignIn(store.User{Email: email})
	if err != nil {
		return "", err
	}

	// Verify the password is valid
	err = auth.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// Get auth token
	token, err := auth.CreateToken(user.ID, secrets.LoadedSecrets.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
