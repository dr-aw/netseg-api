package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/service"
)

type NetSegmentHandler struct {
	service *service.NetSegmentService
}

func NewNetSegmentHandler(service *service.NetSegmentService) *NetSegmentHandler {
	return &NetSegmentHandler{service: service}
}

func RegisterNetSegmentRoutes(e *echo.Echo, handler *NetSegmentHandler) {
	api := e.Group("/api/v1")
	api.GET("/segments", handler.GetAllNetSegments)
	api.POST("/segments", handler.CreateNetSegment)
	api.PUT("/segments/:id", handler.UpdateNetSegment)
}

func (h *NetSegmentHandler) GetAllNetSegments(c echo.Context) error {
	segments, err := h.service.GetAllNetSegments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch segments", "details": err.Error()})
	}
	return c.JSON(http.StatusOK, segments)
}

func (h *NetSegmentHandler) CreateNetSegment(c echo.Context) error {
	var segment domain.NetSegment
	if err := c.Bind(&segment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input", "details": err.Error()})
	}
	if err := h.service.CreateNetSegment(&segment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create segment", "details": err.Error()})
	}
	return c.JSON(http.StatusCreated, segment)
}

func (h *NetSegmentHandler) UpdateNetSegment(c echo.Context) error {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid segment ID", "details": err.Error()})
	}

	var segment domain.NetSegment
	if err := c.Bind(&segment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input", "details": err.Error()})
	}

	segment.ID = uint(idUint)

	if err := h.service.UpdateNetSegment(&segment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update segment", "details": err.Error()})
	}

	return c.JSON(http.StatusOK, segment)
}
