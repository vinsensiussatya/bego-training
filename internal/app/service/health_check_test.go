package service

import (
	"context"
	"os"
	"testing"

	"github.com/vinsensiussatya/bego-training/config"
	"github.com/vinsensiussatya/bego-training/internal/pkg/response"
	"github.com/vinsensiussatya/bego-training/internal/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckService_Ping(t *testing.T) {
	t.Parallel()

	_ = os.Setenv("BEGO_ENV", "test")
	util.GoToProjectDir()
	config.InitConfig()

	// fields successful
	db := config.InitDb(config.GetAppConfig().Database)
	redisCache := config.InitRedis(config.GetAppConfig().Redis)

	// fields failed
	ctxError, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		Database *pgxpool.Pool
		Redis    *redis.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   response.PingResponse
	}{
		{
			name: "OK",
			fields: fields{
				Database: db,
				Redis:    redisCache,
			},
			args: args{
				ctx: context.Background(),
			},
			want: response.PingResponse{
				Database: "OK",
				Redis:    "OK",
			},
		},
		{
			name: "ERROR",
			fields: fields{
				Database: db,
				Redis:    redisCache,
			},
			args: args{
				ctx: ctxError,
			},
			want: response.PingResponse{
				Database: "ERROR: context canceled",
				Redis:    "ERROR: context canceled",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewHealthCheckService(tt.fields.Database, tt.fields.Redis)
			res := s.Ping(tt.args.ctx)

			assert.Equal(t, res, tt.want)
		})
	}
}
