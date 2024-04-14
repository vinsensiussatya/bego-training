package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDb(conf DatabaseConfig) *pgxpool.Pool {
	// load configuration database master
	// connStringExample := "postgres://username:password@localhost:5432/database_name"
	masterConnString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	masterConnConfig, err := pgxpool.ParseConfig(masterConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config: %v\n", err)
		os.Exit(1)
	}

	// set timeout for connection
	masterConnConfig.ConnConfig.ConnectTimeout = time.Duration(conf.Timeout) * time.Second

	// create a connection to database master
	conn, err := pgxpool.NewWithConfig(context.Background(), masterConnConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
