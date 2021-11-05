package services

import (
	"errors"

	"github.com/orlovssky/gread/internal/store"
	"github.com/orlovssky/gread/pkg/auth"
)

type UserService struct{}

var userStore store.UserStore

// Create - Create a user
func (s UserService) Create(user store.User) (store.User, error) {
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
// func (s *UserService) Update(body interface{}, userID int) (store.User, error) {
// 	mbody, ok := body.(map[string]interface{})
// 	if !ok {
// 		return store.User{}, errors.New("cannot map body interface")
// 	}

// 	u, err := userStore.GetById(userID)
// 	if err != nil {
// 		return store.User{}, errors.New("cannot get user by id")
// 	}

// 	for k, v := range mbody {
// 		u.k = v
// 	}

// 	user, err := s.UserStore.Update(u)
// 	if err != nil {
// 		return store.User{}, err
// 	}
// 	return user, nil
// }

// Delete - Deletes a user
func (s UserService) Delete(user store.User) error {
	err := userStore.Delete(user)
	if err != nil {
		return err
	}
	return nil
}
