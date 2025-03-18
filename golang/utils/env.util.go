package utils

import "os"

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func IsLocalEnv() bool {
	return GetEnv("APP_ENV", "development") == "local"
}

func IsDevelopmentEnv() bool {
	return GetEnv("APP_ENV", "development") == "development"
}

func IsProductionEnv() bool {
	return GetEnv("APP_ENV", "development") == "production"
}

func AppName() string {
	return GetEnv("APP_NAME", "golang-app")
}
