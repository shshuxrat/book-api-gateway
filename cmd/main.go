package main

import (
	"book-api-gateway/api"
	"book-api-gateway/config"
	"book-api-gateway/pkg/logger"
	"book-api-gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "example_api_gateway")

	gprcClients, _ := services.NewServicesRepo(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	server.Run(cfg.HttpPort)
}
