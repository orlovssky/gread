package store

import "errors"

// User - An app user
type User struct {
	Base
	Username  string `json:"username" gorm:"type:varchar(255);unique;default:null;"`
	Email     string `json:"email" gorm:"type:varchar(255);unique;not null;"`
	Password  string `json:"password" gorm:"not null;"`
	Firstname string `json:"firstname" gorm:"type:varchar(255);default:null;"`
	Lastname  string `json:"lastname" gorm:"type:varchar(255);default:null;"`
}

type UserStore struct{}

// Create - Creates a user
func (s UserStore) Create(user User) (User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Get - Gets a user
func (s UserStore) Get(user User) (User, error) {
	u := User{}
	if err := DB.Table("users").Where("username=?", user.Username).Or("email = ?", user.Email).Take(&u).Error; err != nil {
		if err.Error() == "sql: no rows in result set" {
			return u, errors.New("user does not exist")
		}
		return u, err
	}
	return u, nil
}

// Get - Gets a user by id
func (s UserStore) GetById(userId int) (User, error) {
	u := User{}
	if err := DB.Table("users").Where("id=?", userId).Take(&u).Error; err != nil {
		if err.Error() == "sql: no rows in result set" {
			return u, errors.New("user does not exist")
		}
		return u, err
	}
	return u, nil
}

// Update - Updates a user
func (s UserStore) Update(mbody map[string]interface{}, userId int) error {
	user := User{}
	user.ID = userId

	if err := DB.Model(&user).Updates(mbody).Error; err != nil {
		return err
	}

	return nil

}

// UpdatePassword - Updates a user's password
func (s UserStore) UpdatePassword(password string, userId int) error {
	if err := DB.Table("users").Where("id=?", userId).UpdateColumn("password", password).Error; err != nil {
		return err
	}

	return nil
}

// Delete - Deletes a user
func (s UserStore) Delete(user User) error {
	if err := DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
