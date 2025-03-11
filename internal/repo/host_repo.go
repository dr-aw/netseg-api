package repo

import (
	"gorm.io/gorm"

	"github.com/dr-aw/netseg-api/internal/domain"
)

type HostRepo struct {
	db *gorm.DB
}

func NewHostRepo(db *gorm.DB) *HostRepo {
	return &HostRepo{db: db}
}

func (r *HostRepo) Create(host *domain.Host) error {
	return r.db.Create(host).Error
}

func (r *HostRepo) GetAll() ([]domain.Host, error) {
	var hosts []domain.Host

	err := r.db.Find(&hosts).Error
	if err != nil {
		return nil, err
	}

	return hosts, nil
}

func (r *HostRepo) Update(host *domain.Host) error {
	return r.db.Save(host).Error
}

func (r *HostRepo) CountHostsBySegmentID(segmentID uint) (int, error) {
	var count int64
	err := r.db.Model(&domain.Host{}).Where("segment_id = ?", segmentID).Count(&count).Error
	return int(count), err
}

func (r *HostRepo) GetByIPAddressAndSegment(ip string, segmentID uint) (*domain.Host, error) {
	var host domain.Host
	err := r.db.Where("ip_address = ? AND segment_id = ?", ip, segmentID).First(&host).Error
	if err != nil {
		return nil, err
	}
	return &host, nil
}

func (r *HostRepo) GetByMAC(mac string) (*domain.Host, error) {
	var host domain.Host
	err := r.db.Where("mac = ?", mac).First(&host).Error
	if err != nil {
		return nil, err
	}
	return &host, nil
}
