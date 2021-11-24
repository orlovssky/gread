package services

import (
	"errors"
	"time"

	readability "github.com/go-shiori/go-readability"
	"github.com/orlovssky/gread/internal/store"
	pkgLink "github.com/orlovssky/gread/pkg/link"
)

type LinkService struct{}

var linkStore store.LinkStore

// Create - Create a user
func (s LinkService) Create(ltp store.LinkToParse, userId int) (store.Link, error) {
	pkgLink.Prepare(&ltp)

	if err := pkgLink.Validate(ltp); err != nil {
		return store.Link{}, err
	}

	parsedLink, err := readability.FromURL(ltp.Link, 30*time.Second)
	if err != nil {
		return store.Link{}, err
	}

	if parsedLink.Content == "" {
		return store.Link{}, errors.New("cannot parse content from link")
	}

	link := store.Link{
		UserID:      userId,
		Url:         ltp.Link,
		Title:       parsedLink.Title,
		Author:      parsedLink.Byline,
		Content:     parsedLink.Content,
		TextContent: parsedLink.TextContent,
		Length:      parsedLink.Length,
		Excerpt:     parsedLink.Excerpt,
		SiteName:    parsedLink.SiteName,
		Image:       parsedLink.Image,
		Favicon:     parsedLink.Favicon,
	}
	// Store link
	link, err = linkStore.Create(link)
	if err != nil {
		return link, err
	}

	return link, nil
}

// Delete - Delete a user
func (s LinkService) Get(linkId int, userId int) (store.Link, error) {
	link, err := linkStore.Get(linkId, userId)
	if err != nil {
		return store.Link{}, err
	}

	return link, nil
}

// GetList - get links
func (s LinkService) GetList(userId int) ([]store.Link, error) {
	var link []store.Link

	link, err := linkStore.GetList(userId)
	if err != nil {
		return link, err
	}

	return link, nil
}

// Delete - Delete a user
func (s LinkService) Delete(linkId int, userId int) error {
	link, err := linkStore.Get(linkId, userId)
	if err != nil {
		return err
	}

	err = linkStore.Delete(link)
	if err != nil {
		return err
	}
	return nil
}
