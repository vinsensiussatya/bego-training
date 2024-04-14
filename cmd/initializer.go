package cmd

import (
	"github.com/vinsensiussatya/bego-training/config"
	"github.com/vinsensiussatya/bego-training/internal/app/handler"
	"github.com/vinsensiussatya/bego-training/internal/app/service"
	"github.com/vinsensiussatya/bego-training/internal/server"
)

func initInjections() server.Handler {
	appConfig := config.GetAppConfig()
	db := config.InitDb(appConfig.Database)

	// wiring services
	hcs := service.NewHealthCheckService(db)

	// wiring handlers
	hch := handler.NewHealthCheckHandler(hcs)

	return server.Handler{
		HealthCheckHandler: hch,
	}
}
