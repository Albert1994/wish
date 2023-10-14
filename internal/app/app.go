package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"wish/internal/config"
	"wish/internal/handler"
	"wish/internal/repository"
	"wish/internal/server"
	"wish/internal/service"
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

	a := os.Getenv("APP_ENV")
	fmt.Println(a)

	db, err := repository.NewPostgresDB(cfg)

	if err != nil {
		logger.Error(err)

		return
	}

	repos := repository.NewRepository(db)

	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	fmt.Println("Server started1")
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	fmt.Println("Server started")
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	fmt.Println("Server started")
	logger.Info("Server started")

}
