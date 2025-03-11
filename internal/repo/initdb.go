package repo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dr-aw/netseg-api/config"
	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/logger"
)

func InitDB(cfg *config.Config) *gorm.DB {
	// Set up the log layer (declared in netseg_repo.go)
	log = logger.LogWithLayer("DATABASE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	log.Infof("Attempting to connect to database %s on host %s:%d", cfg.DBName, cfg.DBHost, cfg.DBPort)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}

	log.Info("Successfully connected to database")

	log.Debug("Running AutoMigrate for Host and NetSegment models...")
	err = db.AutoMigrate(&domain.Host{}, &domain.NetSegment{})
	if err != nil {
		log.WithError(err).Fatal("Failed to migrate database")
	}

	log.Info("Database migration completed successfully")
	return db
}