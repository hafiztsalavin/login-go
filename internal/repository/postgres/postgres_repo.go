package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"login-go/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresRepo initialize new PostgresRepo
func NewPostgresRepo(options *config.DatabaseOption) (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stderr, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second / 100000, // Slow SQL threshold
			LogLevel:      logger.Silent,        // Log level
			Colorful:      true,                 // Disable color
		},
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}
