package databases

import (
	"behemoth-pkg/golang/utils"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/truongkma/gormzerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase establishes a connection to PostgreSQL using GORM.
func ConnectDatabase() *gorm.DB {
	logger := gormzerolog.NewLogger(gormzerolog.Config{
		SlowThreshold:        time.Second,
		ParameterizedQueries: utils.GetEnv("LOGGER_SQL_PARAMETERIZED_QUERIES", "false") == "true",
	})

	db, err := gorm.Open(
		postgres.Open(Dsn()), // ✅ Now calling from `database` package, avoiding circular import
		&gorm.Config{Logger: logger},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("❌ Failed to connect to database")
		os.Exit(1)
	}

	log.Info().Msg("✅ Connected to PostgreSQL")

	ConfigureDatabase(db)

	return db
}

// Close gracefully closes the database connection.
func Close(db *gorm.DB) error {
	if db != nil {
		sqlDb, err := db.DB()
		if err != nil {
			return err
		}
		log.Info().Msg("✅ Closing database connection")
		return sqlDb.Close()
	}
	return nil
}
