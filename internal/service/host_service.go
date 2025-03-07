package service

import (
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
