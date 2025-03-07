package domain

import "time"

type Host struct {
	ID        uint   `gorm:"primaryKey"`
	IPAddress string `gorm:"not null"`
	MAC       string `gorm:"unique;not null"`
	Status    string `gorm:"not null"`
	SegmentID uint   `gorm:"not null;index"`
	UpdatedAt time.Time
}
