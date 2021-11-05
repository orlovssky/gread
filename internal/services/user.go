package services

import (
	"errors"

	"github.com/orlovssky/gread/internal/store"
	"github.com/orlovssky/gread/pkg/auth"
)

type UserService struct {
	UserStore store.UserStoreInterface
}

type UserServiceInterface interface {
	Create(user store.User) (store.User, error)
	Get(user store.User) (store.User, error)
	Update(body interface{}, userID int) (store.User, error)
	Delete(user store.User) error
}

var UserServiceInstance UserServiceInterface = &UserService{}

// Create - Create a user
func (s *UserService) Create(user store.User) (store.User, error) {
	// Check if user already exists
	u, _ := s.UserStore.Get(user)
	if u.ID > 0 {
		return store.User{}, errors.New("this user already exists")
	}

	// Hash password before we store it
	pass, err := auth.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = string(pass)

	// Store user
	user, err = s.UserStore.Create(user)
	if err != nil {
		return user, err
	}
	// Remove passowrd before returning
	user.Password = ""
	return user, nil
}

// Get - Get a user
func (s *UserService) Get(user store.User) (store.User, error) {
	user, err := s.UserStore.Get(user)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Update - Updates a user
func (s *UserService) Update(body interface{}, userID int) (store.User, error) {
	mbody, ok := body.(map[string]interface{})
	if !ok {
		return store.User{}, errors.New("cannot map body interface")
	}

	u, err := s.UserStore.GetById(userID)
	if err != nil {
		return store.User{}, errors.New("cannot get user by id")
	}

	for k, v := range mbody {
		u.k = v
	}

	user, err := s.UserStore.Update(u)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Delete - Deletes a user
func (s *UserService) Delete(user store.User) error {
	err := s.UserStore.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
