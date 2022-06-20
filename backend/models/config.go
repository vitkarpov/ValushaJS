package models

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BindAddr         string
	DatabaseName     string
	DatabasePassword string
	DatabaseUsername string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		BindAddr:         ":8080",
		DatabaseUsername: os.Getenv("POSTGRES_USER"),
		DatabasePassword: os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName:     os.Getenv("POSTGRES_DB"),
	}
}
