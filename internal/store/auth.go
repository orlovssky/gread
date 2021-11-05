package store

import "errors"

type AuthStore struct{}

// Login - Returns user of provided email
func (a AuthStore) SignIn(user User) (User, error) {
	u := User{}
	if err := DB.Table("users").Where("email = ?", user.Email).Take(&u).Error; err != nil {
		if err.Error() == "record not found" {
			return u, errors.New("user does not exist")
		}
		return u, err
	}
	return u, nil
}
