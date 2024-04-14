package repository

import (
	"os"
	"testing"

	"github.com/vinsensiussatya/bego-training/config"
	"github.com/vinsensiussatya/bego-training/internal/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ConnDb *pgxpool.Pool

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

	m.Run()
}
