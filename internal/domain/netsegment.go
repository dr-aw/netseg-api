package domain

import "time"

type NetSegment struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CIDR      string `gorm:"not null"`
	DHCP      bool
	MaxHosts  int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Host struct {
	ID        uint   `gorm:"primaryKey"`
	IPAddress string `gorm:"not null"`
	MAC       string `gorm:"not null"`
	Status    string `gorm:"not null"`
	SegmentID uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
