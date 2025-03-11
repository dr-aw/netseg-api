package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/service"
	"github.com/dr-aw/netseg-api/tests/mocks" // mocks/mock_repo.go
)

// setupTest creates a new NetSegmentService instance and its dependencies for testing.
func setupTest() (*service.NetSegmentService, *mocks.MockNetSegmentBaseRepo, *mocks.MockNetSegmentQueryRepo, *mocks.MockHostQueryRepo) {
	mockNetSegBaseRepo := new(mocks.MockNetSegmentBaseRepo)
	mockNetSegQueryRepo := new(mocks.MockNetSegmentQueryRepo)
	mockHostQueryRepo := new(mocks.MockHostQueryRepo)

	netSegService := service.NewNetSegmentService(mockNetSegBaseRepo, mockNetSegQueryRepo, mockHostQueryRepo)

	return netSegService, mockNetSegBaseRepo, mockNetSegQueryRepo, mockHostQueryRepo
}

// This is a unit test function that verifies if creating a network segment works correctly.
func TestCreateNetSegment_Success(t *testing.T) {
	netSegService, mockNetSegBaseRepo, mockNetSegQueryRepo, _ := setupTest()

	segment := &domain.NetSegment{
		ID:       1,
		Name:     "New Segment",
		CIDR:     "192.168.2.0/24",
		MaxHosts: 20,
	}

	// Check if the segment already exists
	mockNetSegQueryRepo.On("GetByCIDR", segment.CIDR).Return((*domain.NetSegment)(nil), nil)

	// Create the segment
	mockNetSegBaseRepo.On("Create", segment).Return(nil)

	err := netSegService.CreateNetSegment(segment)

	assert.NoError(t, err)
	mockNetSegBaseRepo.AssertExpectations(t)
	mockNetSegQueryRepo.AssertExpectations(t)
}

// This is a unit test function that verifies if updating a network segment works correctly.
func TestUpdateNetSegment_Success(t *testing.T) {
	// Initialize the service and mocks
	netSegService, mockNetSegBaseRepo, mockNetSegQueryRepo, mockHostQueryRepo := setupTest()

	segment := &domain.NetSegment{
		ID:       1,
		Name:     "Test Segment",
		CIDR:     "192.168.1.0/18",
		MaxHosts: 10,
	}

	mockHostQueryRepo.On("CountHostsBySegmentID", segment.ID).Return(5, nil)
	mockNetSegBaseRepo.On("Update", segment).Return(nil)
	mockNetSegQueryRepo.On("GetByCIDR", segment.CIDR).Return((*domain.NetSegment)(nil), nil)

	err := netSegService.UpdateNetSegment(segment)

	assert.NoError(t, err)
	mockHostQueryRepo.AssertExpectations(t)
	mockNetSegBaseRepo.AssertExpectations(t)
	mockNetSegQueryRepo.AssertExpectations(t)
}

// This is a unit test function that verifies if updating a network segment with a CIDR that already exists fails.
func TestUpdateNetSegment_CIDRAlreadyExists(t *testing.T) {
	// Initialize the service and mocks
	netSegService, mockNetSegBaseRepo, mockNetSegQueryRepo, mockHostQueryRepo := setupTest()

	segment := &domain.NetSegment{
		ID:       1,
		Name:     "Test Segment",
		CIDR:     "192.168.1.0/24",
		MaxHosts: 10,
	}

	existingSegment := &domain.NetSegment{
		ID:   2,
		CIDR: "192.168.1.0/24",
	}

	mockHostQueryRepo.On("CountHostsBySegmentID", segment.ID).Return(5, nil)
	mockNetSegQueryRepo.On("GetByCIDR", segment.CIDR).Return(existingSegment, nil)

	err := netSegService.UpdateNetSegment(segment)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "CIDR 192.168.1.0/24 already exists")

	mockNetSegBaseRepo.AssertNotCalled(t, "Update")
}

// This is a unit test function that verifies if getting all network segments works correctly.
func TestGetAllNetSegments_Success(t *testing.T) {
	netSegService, _, mockNetSegQueryRepo, _ := setupTest()

	segments := []domain.NetSegment{
		{ID: 1, Name: "Segment 1", CIDR: "192.168.1.0/24", MaxHosts: 50},
		{ID: 2, Name: "Segment 2", CIDR: "192.168.2.0/24", MaxHosts: 30},
	}

	// 🔹 Ожидаем, что `GetAll()` вернёт два сегмента
	mockNetSegQueryRepo.On("GetAll").Return(segments, nil)

	result, err := netSegService.GetAllNetSegments()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Segment 1", result[0].Name)
	assert.Equal(t, "Segment 2", result[1].Name)

	mockNetSegQueryRepo.AssertExpectations(t)
}

// This is a unit test function that verifies if getting all network segments fails.
func TestGetAllNetSegments_Fail(t *testing.T) {
	netSegService, _, mockNetSegQueryRepo, _ := setupTest()

	// Create a mock that returns an error
	mockNetSegQueryRepo.On("GetAll").Return(nil, fmt.Errorf("repository is not initialized"))

	segments, err := netSegService.GetAllNetSegments()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "repository is not initialized")
	assert.Nil(t, segments)

	mockNetSegQueryRepo.AssertExpectations(t)
}
