package service

import (
	"context"
	"fmt"

	"github.com/vinsensiussatya/bego-training/internal/pkg/response"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type IHealthCheckService interface {
	Ping(ctx context.Context) response.PingResponse
}

type HealthCheckService struct {
	Database *pgxpool.Pool
	Redis    *redis.Client
}

func NewHealthCheckService(database *pgxpool.Pool, redis *redis.Client) IHealthCheckService {
	return &HealthCheckService{
		Database: database,
		Redis:    redis,
	}
}

func (s *HealthCheckService) Ping(ctx context.Context) response.PingResponse {
	dbCh := make(chan string)
	redisCh := make(chan string)
	defer close(dbCh)
	defer close(redisCh)

	go func() {
		dbCh <- checkDB(ctx, s.Database)
	}()

	go func() {
		redisCh <- checkRedis(ctx, s.Redis)
	}()

	return response.PingResponse{
		Database: <-dbCh,
		Redis:    <-redisCh,
	}
}

func checkDB(ctx context.Context, conn *pgxpool.Pool) string {
	err := conn.Ping(ctx)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	return "OK"
}

func checkRedis(ctx context.Context, redisCache *redis.Client) string {
	err := redisCache.Ping(ctx).Err()
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	return "OK"
}
