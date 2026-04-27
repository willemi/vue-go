package config

import (
	"os"

	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Port string
)

func Init() {
	// Load environment variables
	Port = getEnv("PORT", "8080")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}