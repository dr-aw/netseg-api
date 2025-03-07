package repo

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dr-aw/netseg-api/config"
	"github.com/dr-aw/netseg-api/internal/domain"
)

type NetSegmentRepo struct {
	db *gorm.DB
}

func NewNetSegmentRepo(db *gorm.DB) *NetSegmentRepo {
	return &NetSegmentRepo{db: db}
}

func (r *NetSegmentRepo) Create(segment *domain.NetSegment) error {
	return r.db.Create(segment).Error
}

func (r *NetSegmentRepo) GetAll() ([]domain.NetSegment, error) {
	var segments []domain.NetSegment
	err := r.db.Find(segments).Error
	return segments, err
}

func (r *NetSegmentRepo) Update(segment *domain.NetSegment) error {
	return r.db.Save(segment).Error
}

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&domain.NetSegment{}, &domain.Host{})
	return db
}
