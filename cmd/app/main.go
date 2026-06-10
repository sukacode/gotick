package main

import (
	"log"

	"gotick/internal/config"
	"gotick/internal/transport/httpclient"
)

func main() {

	cfg := config.Load()

	client := httpclient.New(
		cfg.BaseURL,
		cfg.APIKey,
	)

	log.Println("base url:", client.BaseURL)
	log.Println("api key loaded:", client.APIKey != "")
}