package repo

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dr-aw/netseg-api/config"
	"github.com/dr-aw/netseg-api/internal/domain"
)

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("DEBUG: Running AutoMigrate...")
	err = db.AutoMigrate(&domain.Host{}, &domain.NetSegment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
