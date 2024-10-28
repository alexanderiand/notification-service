package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/alexanderiand/notification-service/internal/app"
	"github.com/alexanderiand/notification-service/pkg/config"
	"github.com/alexanderiand/notification-service/pkg/logger"
)

func main() {
	println("notification_service")

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("configuration successful initialized")

	if err := logger.InitLogger(cfg); err != nil {
		slog.Warn(err.Error())
	}
	slog.Debug("logger successful initialized")

	ctx := context.Background()

	if err := app.Run(ctx, cfg); err != nil {
		slog.Error(err.Error())
		// app.Stop
	}
}
