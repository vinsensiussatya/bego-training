package config

type AppConfig struct {
	App       App
	Database  DatabaseConfig
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

type BasicAuthConfig struct {
	Username string
	Password string
}
