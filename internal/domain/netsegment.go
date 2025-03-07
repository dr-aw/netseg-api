package domain

import "time"

type NetSegment struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CIDR      string `gorm:"unique;not null"`
	DHCP      bool
	MaxHosts  int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
