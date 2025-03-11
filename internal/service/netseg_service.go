package service

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/repo"
)

type NetSegmentService struct {
	repo     *repo.NetSegmentRepo
	hostRepo *repo.HostRepo
}

func NewNetSegmentService(netSegRepo *repo.NetSegmentRepo, hostRepo *repo.HostRepo) *NetSegmentService {
	return &NetSegmentService{
		repo:     netSegRepo,
		hostRepo: hostRepo,
	}
}


func (s *NetSegmentService) CreateNetSegment(segment *domain.NetSegment) error {
	if err := s.validateSegment(segment); err != nil {
		return err
	}

	return s.repo.Create(segment)
}

func (s *NetSegmentService) GetAllNetSegments() ([]domain.NetSegment, error) {
	if s.repo == nil {
		return nil, fmt.Errorf("repository is not initialized")
	}
	return s.repo.GetAll()
}

func (s *NetSegmentService) UpdateNetSegment(segment *domain.NetSegment) error {

	hostCount, err := s.hostRepo.CountHostsBySegmentID(segment.ID)
	if err != nil {
		return fmt.Errorf("failed to count hosts in segment %d: %v", segment.ID, err)
	}

	if segment.MaxHosts > 0 && segment.MaxHosts < hostCount {
		return fmt.Errorf(
			"cannot update max_hosts to %d: there are already %d hosts in this segment",
			segment.MaxHosts, hostCount,
		)
	}
	if err := s.validateSegment(segment); err != nil {
		return err
	}
	return s.repo.Update(segment)
}

func (s *NetSegmentService) GetSegmentByID(id uint) (*domain.NetSegment, error) {
	return s.repo.GetByID(id)
}

func (s *NetSegmentService) validateSegment(segment *domain.NetSegment) error {
	// Checking CIDR
	_, ipNet, err := net.ParseCIDR(segment.CIDR)
	if err != nil {
		return fmt.Errorf("invalid CIDR format: %s", segment.CIDR)
	}

	maskSize, _ := ipNet.Mask.Size()
	possibleHosts := (1 << (32 - maskSize)) - 2 // 2 hosts always used by network and broadcast
	log.Printf("Possible hosts in subnet: %d", possibleHosts)

	// Checking max_hosts
	if segment.MaxHosts <= 0 {
		return errors.New("max_hosts must be greater than 0")
	}
	if segment.MaxHosts > possibleHosts {
		return fmt.Errorf("max_hosts (%d) exceeds available IP addresses in subnet (%d)", segment.MaxHosts, possibleHosts)
	}

	// CIDR is unique
	existingSegment, err := s.repo.GetByCIDR(segment.CIDR)
	if err == nil && existingSegment != nil && existingSegment.ID != segment.ID {
		return fmt.Errorf("CIDR %s already exists", segment.CIDR)
	}

	return nil
}
