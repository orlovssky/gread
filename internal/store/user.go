package store

import (
	"errors"

	"gorm.io/gorm"
)

// User - An app user
type User struct {
	Base
	Username  string `json:"username" gorm:"type:varchar(255);unique;default:null;"`
	Email     string `json:"email" gorm:"type:varchar(255);unique;not null;" validate:"required,email"`
	Password  string `json:"password" gorm:"not null;" validate:"required,min=8,max=30"`
	Firstname string `json:"firstname" gorm:"type:varchar(255);default:null;"`
	Lastname  string `json:"lastname" gorm:"type:varchar(255);default:null;"`
}

type UserStore struct {
	DB *gorm.DB
}

type UserStoreInterface interface {
	Create(user User) (User, error)
	Get(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) error
}

var UserStoreInstance UserStoreInterface = &UserStore{}

// Create - Creates a user
func (s *UserStore) Create(user User) (User, error) {
	if err := s.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Get - Gets a user
func (s *UserStore) Get(user User) (User, error) {
	u := User{}
	if err := s.DB.Table("users").Where("username=?", user.Username).Or("email = ?", user.Email).Take(&u).Error; err != nil {
		if err.Error() == "sql: no rows in result set" {
			return u, errors.New("user does not exist")
		}
		return u, err
	}
	return u, nil
}

// Update - Updates a user
func (s *UserStore) Update(user User) (User, error) {
	if err := s.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Delete - Deletes a user
func (s *UserStore) Delete(user User) error {
	if err := s.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
