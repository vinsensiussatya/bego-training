package middleware

import (
	"net/http"

	"github.com/vinsensiussatya/bego-training/config"
	myhttp "github.com/vinsensiussatya/bego-training/internal/pkg/http"
)

func BasicAuth(next http.Handler) http.Handler {
	username := config.GetAppConfig().BasicAuth.Username
	password := config.GetAppConfig().BasicAuth.Password

	cw := myhttp.CustomWriter{}
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := myhttp.ErrUnauthorized()
		user, pass, ok := r.BasicAuth()
		if !ok {
			cw.WriteError(w, err)
			return
		}

		isValid := (user == username) && (pass == password)
		if !isValid {
			cw.WriteError(w, err)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
