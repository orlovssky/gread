package user

import (
	"errors"
	"html"
	"strings"

	"github.com/badoux/checkmail"

	"github.com/orlovssky/gread/internal/store"
)

func Prepare(u *store.User) {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)
}

func Validate(u store.User) error {
	if u.Email == "" {
		return errors.New("required email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email")
	}

	if u.Password == "" {
		return errors.New("required password")
	}
	if len(u.Password) < 5 {
		return errors.New("password must have at least 5 characters")
	}
	return nil
}

func ValidateForUpdate(mbody map[string]interface{}) error {
	for k, v := range mbody {
		if strings.ToLower(k) == "email" {
			if email, ok := v.(string); ok {
				if v == "" {
					return errors.New("required email")
				}
				if err := checkmail.ValidateFormat(email); err != nil {
					return errors.New("invalid email")
				}
				var u store.User
				if store.DB.Table("users").Where("email = ?", v).First(&u).Error == nil {
					return errors.New("email already exist")
				}
			}
		}
		if strings.ToLower(k) == "password" {
			return errors.New("password is not allowed")
		}
	}

	return nil
}

func ValidatePassword(mbody map[string]interface{}) (string, error) {
	for k, v := range mbody {
		if strings.ToLower(k) == "password" {
			if _, ok := v.(float64); ok {
				return "", errors.New("password's type should be string")
			}
			if pass, ok := v.(string); ok {
				if pass == "" {
					return "", errors.New("required password")
				}
				if len(pass) < 5 {
					return "", errors.New("password must have at least 5 characters")
				}

				return pass, nil
			}
		}
	}

	return "", errors.New("required password")
}
