package domain

import (
	"errors"
	"net"
	"time"
)

type Host struct {
	ID        uint   `gorm:"primaryKey"`
	IPAddress string `gorm:"not null"`
	MAC       string `gorm:"unique;not null"`
	Status    string `gorm:"not null"`
	SegmentID uint   `gorm:"not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func isIPInSubnet(ip, cidr string) (bool, error) {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return false, errors.New("invalid IP address")
	}

	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false, errors.New("invalid CIDR notation")
	}

	return subnet.Contains(ipAddr), nil
}

func (h *Host) Validate(segmentCIDR string, existingIPs []string) error { // Simple error (only one)
	if !isValidIP(h.IPAddress) {
		return errors.New("invalid IP address format")
	}

	// Checking is IP in subnet
	inSubnet, err := isIPInSubnet(h.IPAddress, segmentCIDR)
	if err != nil {
		return err
	}
	if !inSubnet {
		return errors.New("IP address is not within the segment's subnet")
	}

	// Is IP unique
	for _, ip := range existingIPs {
		if ip == h.IPAddress {
			return errors.New("IP address already exists in this segment")
		}
	}

	return nil
}
