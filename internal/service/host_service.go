package service

import (
	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/repo"
)

type HostService struct {
	repo *repo.HostRepo
}

func NewHostService(repo *repo.HostRepo) *HostService {
	return &HostService{repo: repo}
}

func (s *HostService) CreateHost(host *domain.Host) error {
	return s.repo.Create(host)
}

func (s *HostService) GetAllHosts() ([]domain.Host, error) {
	return s.repo.GetAll()
}

func (s *HostService) UpdateHost(host *domain.Host) error {
	return s.repo.Update(host)
}
