package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Env     string `mapstructure:"ENV"`
	GinMode string `mapstructure:"GIN_MODE" default:"release"`
	Port    string `mapstructure:"PORT"`
}

type Database struct {
	Uri string `mapstructure:"connectionURI"`
}

type Config struct {
	App      App
	Database Database
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	var appConfig App
	var databaseConfig Database

	appConfig.Env = os.Getenv("ENV")
	appConfig.GinMode = os.Getenv("GIN_MODE")
	appConfig.Port = os.Getenv("PORT")

	databaseConfig.Uri = os.Getenv("connectionURI")

	return &Config{
		App:      appConfig,
		Database: databaseConfig,
	}
}
