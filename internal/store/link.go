package store

import (
	"errors"
)

// User - An app user
type Link struct {
	Base
	UserID      int    `json:"user_id" gorm:"type:serial;"`
	Title       string `json:"title" gorm:"type:varchar(255);"`
	Author      string `json:"author" gorm:"type:varchar(255);"`
	Content     string `json:"content" gorm:"type:text;"`
	TextContent string `json:"text_content" gorm:"type:text;"`
	Length      int    `json:"length" gorm:"type:int64;"`
	Excerpt     string `json:"excerpt" gorm:"type:text;"`
	SiteName    string `json:"site_name" gorm:"type:varchar(255);"`
	Image       string `json:"image" gorm:"type:text;"`
	Favicon     string `json:"favicon" gorm:"type:text;"`
}
type LinkToParse struct {
	Link string `json:"link"`
}

type LinkStore struct{}

// Create - Creates a link
func (s LinkStore) Create(link Link) (Link, error) {
	if err := DB.Create(&link).Error; err != nil {
		return link, err
	}
	return link, nil
}

// GetList - Gets links
func (s LinkStore) GetList(userId int) ([]Link, error) {
	var l []Link

	if err := DB.Where("user_id = ?", userId).Find(&l).Error; err != nil {
		return l, err
	}
	return l, nil
}

// Get - Gets a link by userId
func (s LinkStore) Get(linkId int, userId int) (Link, error) {
	l := Link{}
	if err := DB.Table("links").Where("id=? AND user_id=?", linkId, userId).Take(&l).Error; err != nil {
		if err.Error() == "record not found" {
			return l, errors.New("link does not exist")
		}
		return l, err
	}
	return l, nil
}

// Delete - Deletes a link
func (s LinkStore) Delete(link Link) error {
	if err := DB.Delete(&link).Error; err != nil {
		return err
	}
	return nil
}
