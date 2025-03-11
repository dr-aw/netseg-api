package service

import (
	"fmt"
	"net"
	"regexp"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/repo"
)

type HostService struct {
	baseRepo     repo.HostBaseRepository
	queryRepo    repo.HostQueryRepository
	segQueryRepo repo.NetSegmentQueryRepository
	//netSegService	*NetSegmentService
}

func NewHostService(baseRepo repo.HostBaseRepository, queryRepo repo.HostQueryRepository, segQueryRepo repo.NetSegmentQueryRepository) *HostService {
	return &HostService{
		baseRepo:     baseRepo,
		queryRepo:    queryRepo,
		segQueryRepo: segQueryRepo,
		//netSegService: netSegService,
	}
}

func (s *HostService) CreateHost(host *domain.Host) error {
	if err := s.validateHost(host); err != nil {
		return err
	}

	return s.baseRepo.Create(host)
}

func (s *HostService) GetAllHosts() ([]domain.Host, error) {
	if s.queryRepo == nil {
		return nil, fmt.Errorf("repository is not initialized")
	}
	return s.queryRepo.GetAll()
}

// func (s *HostService) GetSegmentByID(id uint) (*domain.NetSegment, error) {
// 	return s.netSegService.GetSegmentByID(id)
// }

func (s *HostService) UpdateHost(host *domain.Host) error {
	if err := s.validateHost(host); err != nil {
		return err
	}
	return s.baseRepo.Update(host)
}

func (s *HostService) validateHost(host *domain.Host) error {
	// Getting segment by ID that host belongs to
	segment, err := s.segQueryRepo.GetByID(host.SegmentID)
	if err != nil {
		return fmt.Errorf("segment with ID %d not found", host.SegmentID)
	}

	// Check if segment has reached max hosts
	hostCount, err := s.queryRepo.CountHostsBySegmentID(host.SegmentID)
	if err != nil {
		return fmt.Errorf("failed to count hosts in segment %d: %v", host.SegmentID, err)
	}
	if segment.MaxHosts > 0 && hostCount >= segment.MaxHosts {
		return fmt.Errorf("cannot add host: segment %s has reached max_hosts limit (%d)", segment.CIDR, segment.MaxHosts)
	}

	// Is IP valid
	if net.ParseIP(host.IPAddress) == nil {
		return fmt.Errorf("invalid IP address format: %s", host.IPAddress)
	}

	// Is IP in subnet
	_, ipNet, _ := net.ParseCIDR(segment.CIDR)
	if !ipNet.Contains(net.ParseIP(host.IPAddress)) {
		return fmt.Errorf("IP %s is outside the segment subnet %s", host.IPAddress, segment.CIDR)
	}

	// mac format
	macRegex := regexp.MustCompile(`^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$`)
	if !macRegex.MatchString(host.MAC) {
		return fmt.Errorf("invalid MAC address format: %s", host.MAC)
	}

	// Checking if IP address is already exists in the segment
	existingHost, err := s.queryRepo.GetByIPAddressAndSegment(host.IPAddress, host.SegmentID)
	if err == nil && existingHost != nil && existingHost.ID != host.ID {
		return fmt.Errorf("IP address %s is already in use in segment %s", host.IPAddress, segment.CIDR)
	}

	// Checking if MAC address is already exists
	existingHost, err = s.queryRepo.GetByMAC(host.MAC)
	if err == nil && existingHost != nil && existingHost.ID != host.ID {
		return fmt.Errorf("MAC address %s is already in use", host.MAC)
	}

	return nil
}

// Method for host_handler
func (s *HostService) GetSegmentByID(id uint) (*domain.NetSegment, error) {
	return s.segQueryRepo.GetByID(id) // ✅ Вызываем напрямую `GetByID()`
}
