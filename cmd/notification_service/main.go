package main

import (
	"context"
	"fmt"
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

	// run service
	slog.Info(fmt.Sprintf("The %s@%s, with env: %s started",
		cfg.Service.Name,
		cfg.Service.Version,
		cfg.Env))

	// context with cancel for stopping the service if happened internal error
	ctx, cancel := context.WithCancel(context.Background())

	if err := app.Run(ctx, cfg); err != nil {
		slog.Error(err.Error())
		// TODO: implement advanced error handling
		cancel()
		return
	}
}
