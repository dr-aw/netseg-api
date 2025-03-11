package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/logger"
	"github.com/dr-aw/netseg-api/internal/service"
)

type NetSegmentHandler struct {
	service *service.NetSegmentService
}

func NewNetSegmentHandler(service *service.NetSegmentService) *NetSegmentHandler {
	return &NetSegmentHandler{service: service}
}

// Setting log layer
var log *logrus.Entry = logger.LogWithLayer("HANDLER")

func (h *NetSegmentHandler) GetAllNetSegments(c echo.Context) error {
	log.Info("GET /segments - Fetching all network segments")
	segments, err := h.service.GetAllNetSegments()
	if err != nil {
		log.WithError(err).Error("Failed to fetch network segments")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch segments", "details": err.Error()})
	}
	log.Infof("Successfully retrieved %d segments", len(segments))
	return c.JSON(http.StatusOK, segments)
}

func (h *NetSegmentHandler) CreateNetSegment(c echo.Context) error {
	log.Info("POST /segments - Creating new network segment...")
	var segment domain.NetSegment
	if err := c.Bind(&segment); err != nil {
		log.WithError(err).Warn("Invalid input data for network segment")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input", "details": err.Error()})
	}
	if err := h.service.CreateNetSegment(&segment); err != nil {
		log.WithError(err).Warn("Failed to create network segment")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create segment", "details": err.Error()})
	}
	log.Infof("Successfully created segment %s with CIDR %s", segment.Name, segment.CIDR)
	return c.JSON(http.StatusCreated, segment)
}

func (h *NetSegmentHandler) UpdateNetSegment(c echo.Context) error {
	id := c.Param("id")
	log.Infof("PUT /segments/:%s - Updating new network segment...", id)
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.WithError(err).Warnf("Invalid segment ID: %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid segment ID", "details": err.Error()})
	}

	var segment domain.NetSegment
	if err := c.Bind(&segment); err != nil {
		log.WithError(err).Warn("Invalid input data for network segment")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input", "details": err.Error()})
	}

	segment.ID = uint(idUint)

	if err := h.service.UpdateNetSegment(&segment); err != nil {
		log.WithError(err).Warnf("Failed to update network segment")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update segment", "details": err.Error()})
	}

	log.Infof("Successfully updated segment %s with ID %b and CIDR %s", segment.Name, idUint, segment.CIDR)
	return c.JSON(http.StatusOK, segment)
}
