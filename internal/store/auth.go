package store

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AuthStore struct {
	DB *gorm.DB
}

type AuthStoreInterface interface {
	Login(user User) (User, error)
}

var AuthStoreInstance AuthStoreInterface = &AuthStore{}

// Login - Returns user of provided email
func (a *AuthStore) Login(user User) (User, error) {
	u := User{}
	if err := DB.Table("users").Where("email = ?", user.Email).Take(&u).Error; err != nil {
		fmt.Println(err)
		return u, errors.New("user does not exist")
	}
	return u, nil
}
