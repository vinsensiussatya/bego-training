package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var appCfg AppConfig

func InitConfig() {
	readConfig()
	appCfg = loadConfig()
}

func GetAppConfig() AppConfig {
	return appCfg
}

// if enabled read config from vault, then set to env
func readConfig() {
	// local only, using local env file; on server should directly read from the ENV
	env := os.Getenv("BEGO_ENV")
	if env == "test" {
		_ = godotenv.Load("./params/.env.test")
	} else {
		_ = godotenv.Load("./params/.env")
	}

	// read config from vault when enabled
	isVaultEnabled, err := strconv.ParseBool(os.Getenv("VAULT_ENABLED"))
	if err != nil {
		log.Fatal().Msgf("Fail value VAULT_ENABLED: %v", err)
	}

	if isVaultEnabled {
		config := vault.DefaultConfig()
		config.Address = getStringOrPanic("VAULT_ADDR")
		client, err := vault.NewClient(config)
		if err != nil {
			log.Fatal().Msgf("unable to initialize Vault client: %v", err)
		}
		client.SetToken(getStringOrPanic("VAULT_TOKEN"))
		client.SetNamespace(getStringOrPanic("VAULT_NAMESPACE"))

		secrets, err := client.KVv2(getStringOrPanic("VAULT_KV_PATH")).Get(context.Background(), getStringOrPanic("VAULT_KV_STORAGE_PATH"))
		if err != nil {
			log.Fatal().Msgf("unable to get secrets from Vault: %v | KV Path: %s, Storage: %s", err, getStringOrPanic("VAULT_KV_PATH"), getStringOrPanic("VAULT_KV_STORAGE_PATH"))
		}

		for k, v := range secrets.Data {
			log.Debug().Msgf("config: %s = %s", k, v)
			err = os.Setenv(k, fmt.Sprint(v))
			if err != nil {
				log.Fatal().Msgf("unable to set env: %v", err)
			}
		}
	}
}

func loadConfig() AppConfig {
	app := App{
		Host: getStringOrPanic("APP_HOST"),
		Port: getIntOrPanic("APP_PORT"),
		Name: getStringOrPanic("APP_NAME"),
	}

	db := DatabaseConfig{
		Username: getStringOrPanic("DB_USERNAME"),
		Password: getStringOrPanic("DB_PASSWORD"),
		Host:     getStringOrPanic("DB_HOST"),
		Port:     getIntOrPanic("DB_PORT"),
		DbName:   getStringOrPanic("DB_NAME"),
		Timeout:  getIntOrPanic("DB_TIMEOUT"),
	}

	basicAuth := BasicAuthConfig{
		Username: getStringOrPanic("BASIC_AUTH_USERNAME"),
		Password: getStringOrPanic("BASIC_AUTH_PASSWORD"),
	}

	appConfig := AppConfig{
		App:       app,
		Database:  db,
		BasicAuth: basicAuth,
	}

	return appConfig
}

func getStringOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal().Msgf("config: %s is required", key)
	}
	return val
}

func getIntOrPanic(key string) int {
	valStr := os.Getenv(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Fatal().Msgf("Error: %v | config: %s is required", err, key)
	}
	return val
}

func getDurationOrPanic(key string) time.Duration {
	valStr := os.Getenv(key)
	val, err := time.ParseDuration(valStr)
	if err != nil {
		log.Fatal().Msgf("Error: %v | config: %s is required", err, key)
	}
	return val
}

// func getBoolDefault(key string, v bool) bool {
// 	val, err := strconv.ParseBool(os.Getenv(key))
// 	if err != nil {
// 		return v
// 	}
// 	return val
// }

func getBoolOrPanic(key string) bool {
	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		log.Fatal().Msgf("Error: %v | config: %s is required", err, key)
	}
	return val
}
