package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/service"
)

type NetSegmentHandler struct {
	service *service.NetSegmentService
}

func RegisterNetSegmentRoutes(e *echo.Echo, service *service.NetSegmentService) {
	api := e.Group("/api/v1")
	handler := &NetSegmentHandler{service: service}
	api.GET("/segments", handler.GetAllNetSegments)
	api.POST("/segments", handler.CreateNetSegment)
}

func (h *NetSegmentHandler) GetAllNetSegments(c echo.Context) error {
	segments, err := h.service.GetAllNetSegments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch segments"})
	}
	return c.JSON(http.StatusOK, segments)
}

func (h *NetSegmentHandler) CreateNetSegment(c echo.Context) error {
	var segment domain.NetSegment
	if err := c.Bind(&segment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if err := h.service.CreateNetSegment(&segment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create segment"})
	}
	return c.JSON(http.StatusCreated, segment)
}
