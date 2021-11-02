package services

import (
	"gorm.io/gorm"

	"github.com/orlovssky/gread/internal/store"
	"github.com/orlovssky/gread/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB        *gorm.DB
	JWTSecret string
	AuthStore store.AuthStoreInterface
}

type AuthServiceInterface interface {
	Login(email, password string) (string, error)
}

var AuthServiceInstance AuthServiceInterface = &AuthService{}

// Login - Here we checek the user email exists and the passwords match.
// Finally an auth token is created and returned
func (as *AuthService) Login(email, password string) (string, error) {
	// Check user with that email exists
	user, err := as.AuthStore.Login(store.User{Email: email})
	if err != nil {
		return "", err
	}

	// Verify the password is valid
	err = auth.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// Get auth token
	token, err := auth.CreateToken(user.ID, as.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
