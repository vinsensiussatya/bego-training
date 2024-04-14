package server

import (
	"net/http"

	myhttp "github.com/vinsensiussatya/bego-training/internal/pkg/http"
	mymiddleware "github.com/vinsensiussatya/bego-training/internal/pkg/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func router(h Handler) *chi.Mux {
	// prepare middleware
	myHandler := myhttp.NewHTTPHandler()

	allowedOrigins := []string{"*"}

	cors := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-ID"},
	})

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(mymiddleware.ContextLogger)
	r.Use(cors.Handler)

	r.Route("/health", func(r chi.Router) {
		r.Use(mymiddleware.BasicAuth)
		r.Method(http.MethodGet, "/readiness", myHandler(h.HealthCheckHandler.Readiness))
		r.Method(http.MethodGet, "/liveness", myHandler(h.HealthCheckHandler.Liveness))
	})

	init_swagger(r)

	return r
}
