package domain

import "time"

type NetSegment struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CIDR      string `gorm:"column:cidr;type:varchar(18);unique;not null" json:"cidr"`
	DHCP      bool
	MaxHosts  int `gorm:"not null" json:"max_hosts"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
