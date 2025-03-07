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

func RegisterHostRoutes(e *echo.Echo, handler *HostHandler) {
	api := e.Group("/api/v1")
	api.GET("/hosts", handler.GetAllHosts)
	api.POST("/hosts", handler.CreateHost)
	api.PUT("/hosts/:id", handler.UpdateHost)
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

	// Validate IP address with mock
	if err := host.Validate("192.168.1.0/24", nil); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Save host
	if err := h.service.CreateHost(&host); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create host"})
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update host"})
	}
	return c.JSON(http.StatusOK, host)
}
