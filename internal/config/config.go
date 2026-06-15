package config

import "os"

type Config struct {
	BaseURL     string
	CorporateID string
	ClientID    string
	ClientSecret string
}

func Load() *Config {
	return &Config{
		BaseURL:      os.Getenv("TIKET_BASE_URL"),
		CorporateID:  os.Getenv("TIKET_CORPORATE_ID"),
		ClientID:     os.Getenv("TIKET_CLIENT_ID"),
		ClientSecret: os.Getenv("TIKET_CLIENT_SECRET"),
	}
}