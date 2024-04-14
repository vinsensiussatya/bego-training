package server

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/vinsensiussatya/bego-training/docs"
)

func init_swagger(r *chi.Mux) {
	docs.SwaggerInfo.Title = "Swagger Fraud Service API"
	docs.SwaggerInfo.Version = "1.0"
	r.Get("/swagger/*", httpSwagger.Handler())
}
