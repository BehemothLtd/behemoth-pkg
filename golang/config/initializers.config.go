package config

import (
	"os"

	"github.com/BehemothLtd/behemoth-pkg/golang/loggers"
	"github.com/BehemothLtd/behemoth-pkg/golang/utils"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitEnvironment(appName string) {
	LoadEnv()
	InitLogger(appName)
}

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Info().Msg("No .env file found, using default values")
		}
	}
}

func InitLogger(appName string) {
	zerolog.SetGlobalLevel(loggers.GetLogLevel())

	appEnv := utils.GetEnv("APP_ENV", "development")

	if utils.GetEnv("APP_DEBUG", "false") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if utils.IsLocalEnv() || utils.IsDevelopmentEnv() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		writer := &loggers.InfoDebugWriter{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		log.Logger = zerolog.New(writer).With().Str("app", appName).Timestamp().Caller().Logger()
	}
	log.Logger = log.Logger.Hook(loggers.TracingHook{})

	log.Logger.Debug().Str("logLevel", zerolog.GlobalLevel().String()).Str("appEnv", appEnv).Msg("🚀 Logger initialized")
}
