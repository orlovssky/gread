package services

import (
	"errors"

	"github.com/orlovssky/gread/internal/store"
	"github.com/orlovssky/gread/pkg/auth"
	pkgUser "github.com/orlovssky/gread/pkg/user"
)

type UserService struct{}

var userStore store.UserStore

// Create - Create a user
func (s UserService) Create(user store.User) (store.User, error) {
	pkgUser.Prepare(&user)

	if err := pkgUser.Validate(user); err != nil {
		return store.User{}, err
	}

	// Check if user already exists
	u, _ := userStore.Get(user)
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
	user, err = userStore.Create(user)
	if err != nil {
		return user, err
	}
	// Remove passowrd before returning
	user.Password = ""
	return user, nil
}

// Get - Get a user
func (s UserService) Get(user store.User) (store.User, error) {
	user, err := userStore.Get(user)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Get - Get a user
func (s *UserService) GetById(userId int) (store.User, error) {
	user, err := userStore.GetById(userId)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Update - Updates a user
func (s *UserService) Update(body interface{}, userId int) error {
	mbody, ok := body.(map[string]interface{})
	if !ok {
		return errors.New("cannot map body interface")
	}

	if err := pkgUser.ValidateForUpdate(mbody); err != nil {
		return err
	}

	err := userStore.Update(mbody, userId)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePassword - Updates a user's password
func (s *UserService) UpdatePassword(body interface{}, userId int) error {
	mbody, ok := body.(map[string]interface{})
	if !ok {
		return errors.New("cannot map body interface")
	}

	password, err := pkgUser.ValidatePassword(mbody)
	if err != nil {
		return err
	}

	// Hash password before we store it
	pass, err := auth.HashPassword(password)
	if err != nil {
		return err
	}
	password = string(pass)

	err = userStore.UpdatePassword(password, userId)
	if err != nil {
		return err
	}
	return nil
}

// Delete - Deletes a user
func (s UserService) Delete(user store.User) error {
	err := userStore.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
