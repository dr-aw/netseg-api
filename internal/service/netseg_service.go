package service

import (
	"errors"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/logger"
	"github.com/dr-aw/netseg-api/internal/repo"
)

type NetSegmentService struct {
	baseRepo      repo.NetSegmentBaseRepository
	queryRepo     repo.NetSegmentQueryRepository
	hostQueryRepo repo.HostQueryRepository
}

// Setting log layer
var log *logrus.Entry = logger.LogWithLayer("SERVICE")

func NewNetSegmentService(base repo.NetSegmentBaseRepository, query repo.NetSegmentQueryRepository, hostQuery repo.HostQueryRepository) *NetSegmentService {
	return &NetSegmentService{
		baseRepo:      base,
		queryRepo:     query,
		hostQueryRepo: hostQuery,
	}
}

func (s *NetSegmentService) CreateNetSegment(segment *domain.NetSegment) error {
	log.WithField("cidr", segment.CIDR).Info("Creating network segment")
	if err := s.validateSegment(segment); err != nil {
		return err
	}
	log.Infof("Segment created successfully: %s", segment.CIDR)
	return s.baseRepo.Create(segment)
}

func (s *NetSegmentService) GetAllNetSegments() ([]domain.NetSegment, error) {
	log.Info("Getting all network segments")
	if s.queryRepo == nil {
		log.Error("Repository is not initialized")
		return nil, fmt.Errorf("repository is not initialized")
	}
	return s.queryRepo.GetAll()
}

func (s *NetSegmentService) UpdateNetSegment(segment *domain.NetSegment) error {
	log.WithField("id", segment.ID).Info("Updating network segment")
	hostCount, err := s.hostQueryRepo.CountHostsBySegmentID(segment.ID)
	if err != nil {
		log.Warnf("failed to count hosts in segment %d: %v", segment.ID, err)
		return fmt.Errorf("failed to count hosts in segment %d: %v", segment.ID, err)
	}

	if segment.MaxHosts > 0 && segment.MaxHosts < hostCount {
		log.Warnf("too low max_hosts: %d, total hosts: %d", segment.MaxHosts, hostCount)
		return fmt.Errorf(
			"cannot update max_hosts to %d: there are already %d hosts in this segment",
			segment.MaxHosts, hostCount,
		)
	}
	if err := s.validateSegment(segment); err != nil {
		return err
	}
	log.Infof("Segment updated successfully: %s", segment.CIDR)
	return s.baseRepo.Update(segment)
}

func (s *NetSegmentService) GetSegmentByID(id uint) (*domain.NetSegment, error) {
	log.Infof("Getting segment by ID: %d", id)
	// error?
	return s.queryRepo.GetByID(id)
}

func (s *NetSegmentService) validateSegment(segment *domain.NetSegment) error {
	// Checking CIDR
	log.Infof("validating segment: %s...", segment.CIDR)
	_, ipNet, err := net.ParseCIDR(segment.CIDR)
	if err != nil {
		log.Warnf("invalid CIDR format: %s", segment.CIDR)
		return fmt.Errorf("invalid CIDR format: %s", segment.CIDR)
	}

	maskSize, _ := ipNet.Mask.Size()
	possibleHosts := (1 << (32 - maskSize)) - 2 // 2 hosts always used by network and broadcast
	log.Infof("Possible hosts in subnet: %d", possibleHosts)

	// Checking max_hosts
	if segment.MaxHosts <= 0 {
		log.Warnf("max_hosts is not set")
		return errors.New("max_hosts must be greater than 0")
	}
	if segment.MaxHosts > possibleHosts {
		log.Warnf("max_hosts (%d) exceeds available IP addresses in subnet (%d)", segment.MaxHosts, possibleHosts)
		return fmt.Errorf("max_hosts (%d) exceeds available IP addresses in subnet (%d)", segment.MaxHosts, possibleHosts)
	}

	// CIDR is unique
	existingSegment, err := s.queryRepo.GetByCIDR(segment.CIDR)
	if err == nil && existingSegment != nil && existingSegment.ID != segment.ID {
		log.Warnf("Attempt to create duplicate CIDR: %s", segment.CIDR)
		return fmt.Errorf("CIDR %s already exists", segment.CIDR)
	}
	log.Infof("Segment validated successfully: %s", segment.CIDR)
	return nil
}
