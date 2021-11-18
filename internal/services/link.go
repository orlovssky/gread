package services

import (
	"time"

	readability "github.com/go-shiori/go-readability"
	"github.com/orlovssky/gread/internal/store"
	pkgLink "github.com/orlovssky/gread/pkg/link"
)

type LinkService struct{}

var linkStore store.LinkStore

// Create - Create a user
func (s LinkService) Create(link store.LinkToParse, userId int) (store.Link, error) {
	pkgLink.Prepare(&link)

	if err := pkgLink.Validate(link); err != nil {
		return store.Link{}, err
	}

	parsedLink, err := readability.FromURL(link.Link, 30*time.Second)
	if err != nil {
		return store.Link{}, err
	}

	l := store.Link{
		UserID:      userId,
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
	l, err = linkStore.Create(l)
	if err != nil {
		return l, err
	}

	return l, nil
}

// GetList - get links
func (s LinkService) GetList(userId int) ([]store.Link, error) {
	var l []store.Link

	l, err := linkStore.GetList(userId)
	if err != nil {
		return l, err
	}

	return l, nil
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
