package service

import (
	"context"
	"fmt"

	"github.com/vinsensiussatya/bego-training/internal/pkg/response"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IHealthCheckService interface {
	Ping(ctx context.Context) response.PingResponse
}

type HealthCheckService struct {
	Database *pgxpool.Pool
}

func NewHealthCheckService(database *pgxpool.Pool) IHealthCheckService {
	return &HealthCheckService{
		Database: database,
	}
}

func (s *HealthCheckService) Ping(ctx context.Context) response.PingResponse {
	dbCh := make(chan string)
	defer close(dbCh)

	go func() {
		dbCh <- checkDB(ctx, s.Database)
	}()

	return response.PingResponse{
		Database: <-dbCh,
	}
}

func checkDB(ctx context.Context, conn *pgxpool.Pool) string {
	err := conn.Ping(ctx)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	return "OK"
}
