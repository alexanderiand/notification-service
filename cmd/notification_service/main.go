package main

import (
	"log"
	"log/slog"

	"github.com/alexanderiand/notification-service/pkg/config"
)

func main() {
	println("notification_service")

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("configuration successful initialized")

	// TODO: init logger
	_ = cfg

	// TODO: run, error handling
}
