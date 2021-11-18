package link

import (
	"errors"
	"html"
	"strings"

	"net/url"

	"github.com/orlovssky/gread/internal/store"
)

func Prepare(ltp *store.LinkToParse) {
	ltp.Link = html.EscapeString(strings.TrimSpace(ltp.Link))
}

func Validate(ltp store.LinkToParse) error {
	if ltp.Link == "" {
		return errors.New("required link")
	}
	_, err := url.ParseRequestURI(ltp.Link)
	if err != nil {
		return errors.New("invalid link")
	}
	return nil
}
