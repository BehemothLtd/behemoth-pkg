package config

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BehemothLtd/behemoth-pkg/golang/databases"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

const (
	shutdownTimeout = 5 * time.Second
)

// StartServer begins listening and serving HTTP requests
func StartServer(srv *http.Server, appPort string) {
	err := srv.ListenAndServe()
	log.Info().Msg("🚀 Server running at http://localhost" + appPort)

	if err != nil {
		log.Error().Err(err).Msg("❌ Failed to start API server")
		os.Exit(1)
	}
}

// WaitForShutdown gracefully shuts down the server when an interruption signal is received
func WaitForShutdown(srv *http.Server, db *gorm.DB) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown signal received, shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("❌ Failed to gracefully shutdown API server")
	} else {
		log.Info().Msg("✅ Server shutdown complete")
	}

	if err := databases.Close(db); err != nil {
		log.Info().Msg("✅ Database connection closed")
	}

	log.Info().Msg("👋 Bye!")
}
