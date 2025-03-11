package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/dr-aw/netseg-api/internal/domain"
	"github.com/dr-aw/netseg-api/internal/service"
)

type HostHandler struct {
	service *service.HostService
}

func NewHostHandler(service *service.HostService) *HostHandler {
	return &HostHandler{service: service}
}

func (h *HostHandler) GetAllHosts(c echo.Context) error {
	hosts, err := h.service.GetAllHosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch hosts"})
	}
	return c.JSON(http.StatusOK, hosts)
}

func (h *HostHandler) CreateHost(c echo.Context) error {
	var host domain.Host
	if err := c.Bind(&host); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	segment, err := h.service.GetSegmentByID(host.SegmentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Segment not found"})
	}

	// Validate IP address in subnet
	if err := host.Validate(segment.CIDR, nil); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Save host
	if err := h.service.CreateHost(&host); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create host", "details": err.Error()})
	}
	return c.JSON(http.StatusCreated, host)
}

func (h *HostHandler) UpdateHost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid host ID"})
	}

	var host domain.Host
	if err := c.Bind(&host); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	host.ID = uint(id)

	if err := h.service.UpdateHost(&host); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update host", "details": err.Error()})
	}
	return c.JSON(http.StatusOK, host)
}
