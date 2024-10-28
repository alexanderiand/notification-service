package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexanderiand/notification-service/internal/infrastructure/repository"
	"github.com/alexanderiand/notification-service/internal/infrastructure/repository/storage/sqlite"
	"github.com/alexanderiand/notification-service/internal/transport/http/rest/controller"
	"github.com/alexanderiand/notification-service/internal/transport/http/rest/router"
	"github.com/alexanderiand/notification-service/internal/transport/http/rest/server"
	"github.com/alexanderiand/notification-service/internal/usecase"
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

	if err := db.Ping(); err != nil { // ping a db for checking the db connection
		slog.Error(err.Error())
		return err
	}
	slog.Debug("successful connection to database")

	// implement DI
	repo := repository.New(db)
	slog.Debug("successful create a new instance of the notification")
	usecase := usecase.New(repo)
	slog.Debug("successful create a new instance of the usecase - notification usecase")
	controller := controller.New(usecase)
	slog.Debug("successful create a new instance of the controller")

	// create a new instance of the Router eq http.ServeMux
	router := router.New(controller)
	slog.Debug("successful create a new instance of the router")
	router.InitRouter(cfg) // mapping request to controller methods

	// create and start a http server instance
	httpSrv := server.New(cfg, router)

	// Implement graceful shutdown
	sigChan := make(chan os.Signal, 1)

	defer close(sigChan)

	// run the http server into a separate goroutine
	go func() {
		if err := httpSrv.Start(); err != nil {
			slog.Error(err.Error())
		}

		if err := ctx.Err(); err != nil {
			return
		}
	}()

	// listen os signals, like SIGTERM, SIGQUIT, SIGINT
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	sysSignal, sigOk := <-sigChan
	if sigOk {
		slog.Info("service shuting down", "os_signal", sysSignal.String())

		// shuting down the http server
		durTime := time.Second * 3
		ctx, cancel := context.WithTimeout(ctx, durTime)
		defer cancel()

		if err := httpSrv.Shutdown(ctx); err != nil {
			return err
		}
		slog.Debug("the http server successful shuting down")

		if err := db.Close(); err != nil {
			return err
		}
		slog.Debug("the database connection successful closed")
	}

	return nil
}
