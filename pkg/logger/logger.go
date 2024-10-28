package logger

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/alexanderiand/notification-service/pkg/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// errors
var (
	ErrNilStructPointer = errors.New("error, nil struct pointer")
)

// InitLogger depending on the env configuration parameter, sets the log level, and handler type
// If cfg is nil, return ErrNilStructPointer
func InitLogger(cfg *config.Config) error {
	op := "pkg.logger.InitLogger"

	if cfg == nil {
		return fmt.Errorf("%s %w", op, ErrNilStructPointer)
	}

	switch cfg.Env {
	case envLocal:
		log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true, // info about a log record path, line, and column num
		}))
		slog.SetDefault(log)
	case envDev:
		log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
		slog.SetDefault(log)
	case envProd:
		log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		}))
		slog.SetDefault(log)
	default:
		log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		}))
		slog.SetDefault(log)
	}

	return nil
}
