package service

import (
	"fmt"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/repo"
)

type HostService struct {
	repo          *repo.HostRepo
	netSegService *NetSegmentService
}

func NewHostService(repo *repo.HostRepo, netSegService *NetSegmentService) *HostService {
	return &HostService{repo: repo, netSegService: netSegService}
}

func (s *HostService) CreateHost(host *domain.Host) error {
	// Getting segment by ID that host belongs to
	segment, err := s.netSegService.GetSegmentByID(host.SegmentID)
	if err != nil {
		return fmt.Errorf("segment not found: %w", err)
	}

	// Checking max_hosts limit
	hostCount, err := s.repo.CountHostsBySegmentID(host.SegmentID)
	if err != nil {
		return fmt.Errorf("failed to count hosts: %w", err)
	}
	if hostCount >= segment.MaxHosts {
		return fmt.Errorf("cannot add host: segment %s reached max_hosts limit (%d)", segment.CIDR, segment.MaxHosts)
	}

	// Checking if IP address already exists in the segment
	existingHost, err := s.repo.GetByIPAddressAndSegment(host.IPAddress, host.SegmentID)
	if err == nil && existingHost != nil {
		return fmt.Errorf("IP address %s already exists in segment %s", host.IPAddress, segment.CIDR)
	}

	return s.repo.Create(host)
}

func (s *HostService) GetAllHosts() ([]domain.Host, error) {
	return s.repo.GetAll()
}

func (s *HostService) GetSegmentByID(id uint) (*domain.NetSegment, error) {
	return s.netSegService.GetSegmentByID(id)
}

func (s *HostService) UpdateHost(host *domain.Host) error {
	return s.repo.Update(host)
}
