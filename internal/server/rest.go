package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/vinsensiussatya/bego-training/config"
	"github.com/vinsensiussatya/bego-training/internal/app/handler"

	"github.com/rs/zerolog/log"
)

type Server struct {
	handler Handler
}

type Handler struct {
	HealthCheckHandler handler.IHealthCheckHandler
}

func NewServer(h Handler) Server {
	return Server{
		handler: h,
	}
}

func (s Server) StartRestAPI() {
	appConfig := config.GetAppConfig()
	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Info().Msgf("[REST API] Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Error().Msgf("[REST API] Fail to shutting down: %v", err)
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", appConfig.App.Host, appConfig.App.Port)
	srv.Handler = router(s.handler)

	log.Info().Msgf("[API] HTTP serve at %s\n", srv.Addr)
	log.Info().Msgf("[API Docs] Go to http://%s/swagger/", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Error().Msgf("[API] Fail to listen and serve: %v", err)
	}

	<-idleConnectionClosed
	log.Info().Msg("[API] Bye")
}
