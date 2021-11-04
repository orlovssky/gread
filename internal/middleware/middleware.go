package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/orlovssky/gread/api"
	"github.com/orlovssky/gread/pkg/auth"
)

func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id, err := auth.TokenValid(r)
		if err != nil {
			api.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
