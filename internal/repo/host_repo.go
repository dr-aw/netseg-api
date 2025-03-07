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
	err := r.db.Find(hosts).Error
	return hosts, err
}

func (r *HostRepo) Update(host *domain.Host) error {
	return r.db.Save(host).Error
}
