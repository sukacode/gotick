package config

import "os"

type Config struct {
	APIKey string
	APIURL string
}

func Load() *Config {
	return &Config{
		APIKey: os.Getenv("API_KEY"),
		APIURL: os.Getenv("API_URL"),
	}
}