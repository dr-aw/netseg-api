package repo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/logger"
)

type NetSegmentRepo struct {
	db *gorm.DB
}

// Setting log layer
var log *logrus.Entry = logger.LogWithLayer("REPO")

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
	log.Debugf("Attempting to insert segment (CIDR: %s)", segment.CIDR)

	err := r.db.Create(segment).Error
	if err != nil {
		log.WithError(err).Errorf("Failed to insert segment (CIDR: %s) into DB: %v", segment.CIDR, err)
		return err
	}

	log.Infof("Successfully inserted segment: %s", segment.CIDR)
	return nil
}

func (r *NetSegmentRepo) GetAll() ([]domain.NetSegment, error) {
	log.Debug("Fetching all segments from DB")
	var segments []domain.NetSegment
	err := r.db.Find(&segments).Error
	if err != nil {
		log.WithError(err).Error("Failed to fetch all network segments")
		return nil, err
	}
	log.Infof("Retrieved %d segments from DB", len(segments))
	return segments, err
}

func (r *NetSegmentRepo) GetByID(id uint) (*domain.NetSegment, error) {
	log.Debugf("Fetching segment by ID: %d", id)
	var segment domain.NetSegment
	err := r.db.First(&segment, id).Error
	if err != nil {
		log.WithError(err).Warnf("Segment with ID %d not found", id)
		return nil, err
	}
	log.Infof("Found segment (ID: %d, CIDR: %s)", segment.ID, segment.CIDR)
	return &segment, err
}

func (r *NetSegmentRepo) GetByCIDR(cidr string) (*domain.NetSegment, error) {
	log.Debugf("Fetching segment by CIDR: %s", cidr)
	var segment domain.NetSegment
	err := r.db.Where("cidr = ?", cidr).First(&segment).Error
	if err != nil {
		log.WithError(err).Warnf("Segment with CIDR %s not found", cidr)
		return nil, err
	}
	log.Infof("Found segment (ID: %d, CIDR: %s)", segment.ID, segment.CIDR)
	return &segment, err
}

func (r *NetSegmentRepo) Update(segment *domain.NetSegment) error {
	log.Debugf("Attempting to update segment (ID: %d, CIDR: %s)", segment.ID, segment.CIDR)
	err := r.db.Save(segment).Error
	if err != nil {
		log.WithError(err).Errorf("Failed to update segment (ID: %d, CIDR: %s)", segment.ID, segment.CIDR)
		return err
	}
	log.Infof("Successfully updated segment (ID: %d, CIDR: %s)", segment.ID, segment.CIDR)
	return nil
}
