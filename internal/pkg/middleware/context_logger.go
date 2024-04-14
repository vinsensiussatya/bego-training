package middleware

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func ContextLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(log.With().Logger().WithContext(r.Context())) // add sublog context
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
