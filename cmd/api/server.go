package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/pushgate/core/internal/config"
	data "github.com/pushgate/core/internal/data"
	slogger "github.com/pushgate/core/pkg/logger"
	"github.com/pushgate/core/pkg/postgresql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "1.0.0"

type application struct {
	config *config.Config
	logger *slog.Logger
	models data.Models
}

func Run(port int, env string) {
	cfg := config.Load(false)
	cfg.App.Port = port
	cfg.App.Env = env

	initializeApp(cfg)
}

func initializeApp(cfg *config.Config) {
	h := &ContextHandler{slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: false})}
	logger := slogger.NewLogger(h)
	logger.Info("Starting API server", slog.String("version", version))

	db, err := openDB(cfg.Db)
	if err != nil {
		logger.Error("error opening database connection", slog.String("error", err.Error()))
		os.Exit(1)
	}

	defer func(db postgresql.DB) {
		err := db.Close()
		if err != nil {
			logger.Error("error closing database connection", slog.String("error", err.Error()))
		}
	}(db)

	logger.Info("database connection established",
		slog.String("host", cfg.Db.Host),
		slog.String("port", cfg.Db.Port),
		slog.String("database", cfg.Db.Database),
		slog.String("user", cfg.Db.Username),
		slog.String("sslmode", cfg.Db.SSLMode),
	)

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Info("starting server", slog.String("Env", cfg.App.Env), slog.String("Addr", srv.Addr))

	// Graceful shutdown
	shutdownError := make(chan error)

	// Create a new channel to receive OS signals
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		s := <-quit
		app.logger.Info("caught signal", slog.String("signal", s.String()))
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		shutdownError <- srv.Shutdown(ctx)
	}()

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		app.logger.Error("server error", slog.String("error", err.Error()))
	}

	err = <-shutdownError
	if err != nil {
		app.logger.Error("server error", slog.String("error", err.Error()))
	}

	app.logger.Info("stopped server", slog.String("Env", cfg.App.Env), slog.String("Addr", srv.Addr))
}

func openDB(cfg *postgresql.Config) (postgresql.DB, error) {
	db, err := postgresql.New(cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}
