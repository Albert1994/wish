package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wish/internal/config"
	"wish/internal/handler"
	"wish/internal/repository"
	"wish/internal/server"
	"wish/internal/service"
	"wish/pkg/auth"
	"wish/pkg/hash"
	"wish/pkg/logger"
)

// @title Wish API
// @version 1.0
// @description REST API for Wish

// @host localhost:8000
// @BasePath /api/v1/

// @securityDefinitions.apikey StudentsAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

// Run initializes whole application.
func Run(configDir string) {
	cfg, err := config.Init(configDir)
	if err != nil {
		logger.Error(err)

		return
	}

	db, err := repository.NewPostgresDB(cfg)

	if err != nil {
		logger.Error(err)

		return
	}

	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)

		return
	}

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	repos := repository.NewRepository(db)

	services := service.NewServices(service.Deps{
		Repos:        repos,
		TokenManager: tokenManager,
		Hasher:       hasher,
	})
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
