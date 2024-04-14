package config

type AppConfig struct {
	App       App
	Database  DatabaseConfig
	Redis     RedisConfig
	BasicAuth BasicAuthConfig
}

type App struct {
	Host string
	Port int
	Name string
}

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
	Timeout  int
}

type RedisConfig struct {
	Url string
}

type BasicAuthConfig struct {
	Username string
	Password string
}
