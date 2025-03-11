package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/dr-aw/netseg-api/internal/domain"
)

// Mocks for `NetSegmentBaseRepository`
type MockNetSegmentBaseRepo struct {
	mock.Mock
}

func (m *MockNetSegmentBaseRepo) Create(segment *domain.NetSegment) error {
	args := m.Called(segment)
	return args.Error(0)
}

func (m *MockNetSegmentBaseRepo) Update(segment *domain.NetSegment) error {
	args := m.Called(segment)
	return args.Error(0)
}

// Mocks for `NetSegmentQueryRepository`
type MockNetSegmentQueryRepo struct {
	mock.Mock
}

func (m *MockNetSegmentQueryRepo) GetByID(id uint) (*domain.NetSegment, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.NetSegment), args.Error(1)
}

func (m *MockNetSegmentQueryRepo) GetAll() ([]domain.NetSegment, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.NetSegment), args.Error(1)
}

func (m *MockNetSegmentQueryRepo) GetByCIDR(cidr string) (*domain.NetSegment, error) {
	args := m.Called(cidr)
	return args.Get(0).(*domain.NetSegment), args.Error(1)
}

// Mocks for `HostQueryRepository`
type MockHostQueryRepo struct {
	mock.Mock
}

func (m *MockHostQueryRepo) GetByIPAddressAndSegment(ip string, segmentID uint) (*domain.Host, error) {
	args := m.Called(ip, segmentID)
	return args.Get(0).(*domain.Host), args.Error(1)
}

func (m *MockHostQueryRepo) GetByMAC(mac string) (*domain.Host, error) {
	args := m.Called(mac)
	return args.Get(0).(*domain.Host), args.Error(1)
}

func (m *MockHostQueryRepo) GetAll() ([]domain.Host, error) {
	args := m.Called()
	return args.Get(0).([]domain.Host), args.Error(1)
}

func (m *MockHostQueryRepo) CountHostsBySegmentID(segmentID uint) (int, error) {
	args := m.Called(segmentID)
	return args.Int(0), args.Error(1)
}
