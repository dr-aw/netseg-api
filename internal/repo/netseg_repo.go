package repo

import (
	"gorm.io/gorm"

	"github.com/dr-aw/netseg-api/internal/domain"
)

type NetSegmentRepo struct {
	db *gorm.DB
}

// Interface for writing
type NetSegmentBaseRepository interface {
	Create(segment *domain.NetSegment) error
	Update(segment *domain.NetSegment) error
}

// Interface for reading
type NetSegmentQueryRepository interface {
	GetByID(id uint) (*domain.NetSegment, error)
	GetAll() ([]domain.NetSegment, error)
	GetByCIDR(cidr string) (*domain.NetSegment, error)
}


func NewNetSegmentRepo(db *gorm.DB) *NetSegmentRepo {
	return &NetSegmentRepo{db: db}
}

func (r *NetSegmentRepo) Create(segment *domain.NetSegment) error {
	return r.db.Create(segment).Error
}

func (r *NetSegmentRepo) GetAll() ([]domain.NetSegment, error) {
	var segments []domain.NetSegment
	err := r.db.Find(&segments).Error
	return segments, err
}

func (r *NetSegmentRepo) GetByID(id uint) (*domain.NetSegment, error) {
	var segment domain.NetSegment
	err := r.db.First(&segment, id).Error
	return &segment, err
}

func (r *NetSegmentRepo) GetByCIDR(cidr string) (*domain.NetSegment, error) {
	var segment domain.NetSegment
	err := r.db.Where("cidr = ?", cidr).First(&segment).Error
	return &segment, err
}

func (r *NetSegmentRepo) Update(segment *domain.NetSegment) error {
	return r.db.Save(segment).Error
}
