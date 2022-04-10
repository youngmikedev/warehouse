package app

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/youngmikedev/warehouse/internal/auth"
	"github.com/youngmikedev/warehouse/internal/config"
	"github.com/youngmikedev/warehouse/internal/delivery/swagger"
	"github.com/youngmikedev/warehouse/internal/repository"
	"github.com/youngmikedev/warehouse/internal/repository/postgres"
	"github.com/youngmikedev/warehouse/internal/service"
)

func Run() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		logger.Fatal().
			Err(err).
			Msg("Failed init .env file")
		return
	}

	cfg, err := config.Init()
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed init config")
		return
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(cfg.Log.Level))

	// PostrgeSQL
	postgresClient, err := postgres.NewClient(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password,
	)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed init Postgres client")
		return
	}

	repos := repository.NewPostgresRepositories(postgresClient)

	cache := service.NewCacheWithMaxLen(0)
	tokenManager := auth.NewTokenManager(cfg.Token.SigningKey, time.Duration(cfg.Token.UTExpiresAt), time.Duration(cfg.Token.URTExpiresAT))
	hashManager := auth.NewHashManager()
	sl := logger.With().Str("from", "service").Logger()
	services := service.NewServices(repos, cache, tokenManager, hashManager, &sl)

	dl := logger.With().Str("from", "delivery").Logger()
	httpServer, err := swagger.NewServer(services, cfg, &dl)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed init Postgres client")
		return
	}

	// Start server
	go func() {
		if err := httpServer.Serve(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().
				Err(err).
				Msg("Shutting down the http server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := httpServer.Shutdown(); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed shutting down the http server")
	}
	logger.Info().Msg("Http server successful shutdown")

}
