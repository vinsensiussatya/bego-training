package repository

import (
	"context"
	"os"
	"testing"

	"github.com/vinsensiussatya/bego-training/config"
	"github.com/vinsensiussatya/bego-training/internal/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

var ConnDb *pgxpool.Pool
var Rdb *redis.Client

func TestMain(m *testing.M) {
	_ = os.Setenv("BEGO_ENV", "test")
	util.GoToProjectDir()
	config.InitConfig()

	// load config
	util.GoToProjectDir()
	config.InitConfig()

	// connection database
	conn := config.InitDb(config.GetAppConfig().Database)
	ConnDb = conn

	// connection redis
	rdb := config.InitRedis(config.GetAppConfig().Redis)
	Rdb = rdb

	m.Run()

	cleanUpCache(rdb)
}

func cleanUpCache(rdb *redis.Client) {
	err := rdb.FlushAll(context.Background()).Err()
	if err != nil {
		return
	}

	_ = rdb.Close()
}
