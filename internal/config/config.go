package config

import "os"

type Config struct {
	BaseURL string
	APIKey  string
}

func Load() *Config {
	return &Config{
		BaseURL: os.Getenv("TIKET_BASE_URL"),
		APIKey:  os.Getenv("TIKET_API_KEY"),
	}
}