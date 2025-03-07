package repo

import (
	"log"

	"github.com/dr-aw/netseg-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NetSegmentRepo struct {
	db *gorm.DB
}

func (r *NetSegmentRepo) Create(segment *domain.NetSegment) error {
	return nil
}

func (r *NetSegmentRepo) GetAll() ([]domain.NetSegment, error) {
	var segments []domain.NetSegment
	err := r.db.Find(segments).Error
	return segments, err
}

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=password123 dbname=netseg port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	log.Println("db connected successfully")
	// migration
	return db
}
