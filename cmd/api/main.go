package main

import (
	"context"
	"fmt"
	httpx "go.strv.io/net/http"
	"log/slog"
	apiv1 "newsletter_backend_api/cmd/api/v1"
	repo "newsletter_backend_api/repository"
	"newsletter_backend_api/transport/util"

	"newsletter_backend_api/service"

	transport "newsletter_backend_api/transport/api"

	"github.com/jackc/pgx/v5/pgxpool"
)

var version = "v1.0.0"

func main() {
	ctx := context.Background()
	cfg := apiv1.MustLoadConfig()
	//logging currently broken
	//util.SetServerLogLevel(slog.LevelInfo)

	db, err := setupDatabase(ctx, cfg)
	if err != nil {
		slog.Error("initializing database", slog.Any("error", err))
	}
	repository, err := repo.New(db)
	if err != nil {
		slog.Error("initializing repository", slog.Any("error", err))
	}

	controller, err := setupController(
		cfg,
		repository,
	)
	if err != nil {
		slog.Error("initializing controller", slog.Any("error", err))
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	// Initialize the server config.
	serverConfig := httpx.ServerConfig{
		Addr:    addr,
		Handler: controller,
		Hooks:   httpx.ServerHooks{
			// BeforeShutdown: []httpx.ServerHookFunc{
			// 	func(_ context.Context) {
			// 		database.Close()
			// 	},
			// },
		},
		Limits: nil,
		Logger: util.NewServerLogger("httpx.Server").Logger,
	}
	server := httpx.NewServer(&serverConfig)

	slog.Info("starting server", slog.Int("port", cfg.Port))
	if err := server.Run(ctx); err != nil {
		slog.Error("server failed", slog.Any("error", err))
	}
}

func setupDatabase(ctx context.Context, cfg apiv1.Config) (*pgxpool.Pool, error) {
	// Initialize the database connection pool.
	pool, err := pgxpool.New(
		ctx,
		cfg.DatabaseURL,
	)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func setupController(_ apiv1.Config, repository service.Repository) (*transport.Controller, error) {
	svc, err := service.NewService(repository)
	if err != nil {
		return nil, fmt.Errorf("initializing user service: %w", err)
	}

	// Initialize the controller.
	controller, err := transport.NewController(
		svc,
		version,
	)
	if err != nil {
		return nil, fmt.Errorf("initializing controller: %w", err)
	}

	return controller, nil
}
