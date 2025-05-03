package config

import "os"

type Config struct {
	DbConnection string
	DBName       string
}

func NewConfig() *Config {
	return &Config{
		DbConnection: os.Getenv("DB_CONNECTION"),
		DBName:       os.Getenv("DB_NAME"),
	}
}
