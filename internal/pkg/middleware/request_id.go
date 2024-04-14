package middleware

import (
	"net/http"

	myctx "github.com/vinsensiussatya/bego-training/internal/pkg/context"

	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

const xReqID = "X-Request-ID"

func RequestIDAndContextLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(xReqID)
		if reqID == "" {
			reqID = uuid.NewV4().String()
			r.Header.Set(xReqID, reqID) //set request ID to header if not set
		}

		r = r.WithContext(log.With().Str(myctx.RequestID.String(), reqID).Logger().WithContext(r.Context())) // add log with request ID
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
