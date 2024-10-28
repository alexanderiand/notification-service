package app

import (
	"context"
	"log/slog"

	"github.com/alexanderiand/notification-service/internal/infrastructure/repository/storage/sqlite"
	"github.com/alexanderiand/notification-service/pkg/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	// db init
	db, err := sqlite.New(cfg)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	slog.Debug("successful created a new sqlite client")

	if err := db.DB.Ping(); err != nil {
		slog.Error(err.Error())
		return err
	}
	slog.Debug("successful connection to database")

	return nil
}
