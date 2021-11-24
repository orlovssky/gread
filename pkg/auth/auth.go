package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	// "strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/orlovssky/gread/internal/secrets"
	"golang.org/x/crypto/bcrypt"
)

// CreateToken - Creates abd returns a new JTW token for the
// given userID. Token expires after 1 hour
func CreateToken(userID int, apiSecret string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix() //Token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(apiSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}

// TokenValid - Checks if the token passed in the request is
// valid. If not valid an error will reee returned
func TokenValid(r *http.Request) (int, error) {
	tokenString := extractToken(r)
	if tokenString == "" {
		return 0, errors.New("invliad token")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secrets.LoadedSecrets.JWTSecret), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invliad token")
	}
	id, err := extractID(token)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// extractID - Gets the id from claims
func extractID(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, err := strconv.Atoi(fmt.Sprintf("%.f", claims["id"]))
		if err != nil {
			return 0, errors.New("cannot parse id from claims")
		} else {
			return id, nil
		}
	}
	return 0, errors.New("cannot extract id")
}

// extractToken - Gets a token from the req header
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(strings.TrimSpace(bearToken), " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// HashPassword - Returns the bcrypt hash of the password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// HashPassword - Verifies the bcrypt hash of the password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
