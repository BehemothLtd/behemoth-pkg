package databases

import (
	"behemoth-pkg/golang/utils"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Dsn generates the PostgreSQL connection string.
func Dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Ho_Chi_Minh password=%s",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_NAME", "billiard-community-dev"),
		utils.GetEnv("DB_PORT", "5432"),
		getSSLMode(),
		utils.GetEnv("DB_PASSWORD", ""),
	)
}

// ConfigureDatabase sets up connection pool settings for the provided DB instance.
func ConfigureDatabase(db *gorm.DB) {
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Failed to get SQL DB")
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = sqlDb.PingContext(ctxTimeout)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not establish connection to postgres")
	}

	maxOpenConns, _ := strconv.Atoi(utils.GetEnv("DB_MAX_OPEN_CONNS", "20"))
	maxIdleConns, _ := strconv.Atoi(utils.GetEnv("DB_MAX_IDLE_CONNS", "10"))
	connMaxLifetime, _ := strconv.Atoi(utils.GetEnv("DB_CONN_MAX_LIFETIME_HOURS", "1"))

	sqlDb.SetMaxOpenConns(maxOpenConns)
	sqlDb.SetMaxIdleConns(maxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Hour)

	log.Info().Msg("🔧 Database connection settings configured")
}

func getSSLMode() string {
	sslModes := map[string]string{
		"require":     "require",
		"verify-full": "verify-full",
		"verify-ca":   "verify-ca",
		"prefer":      "prefer",
		"allow":       "allow",
		"disable":     "disable",
	}

	sslMode := sslModes[utils.GetEnv("DB_SSL_MODE", "disable")]
	if sslMode == "" {
		sslMode = "disable"
	}

	return sslMode
}
