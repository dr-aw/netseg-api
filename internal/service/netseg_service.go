package service

import (
	"fmt"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/repo"
)

type NetSegmentService struct {
	repo *repo.NetSegmentRepo
}

func NewNetSegmentService(repo *repo.NetSegmentRepo) *NetSegmentService {
	return &NetSegmentService{repo: repo}
}

func (s *NetSegmentService) CreateNetSegment(segment *domain.NetSegment) error {
	return s.repo.Create(segment)
}

func (s *NetSegmentService) GetAllNetSegments() ([]domain.NetSegment, error) {
	if s.repo == nil {
		return nil, fmt.Errorf("repository is not initialized")
	}
	return s.repo.GetAll()
}

func (s *NetSegmentService) UpdateNetSegment(segment *domain.NetSegment) error {
	return s.repo.Update(segment)
}
