package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, netSegHandler *NetSegmentHandler, hostHandler *HostHandler) {
	api := e.Group("/api/v1")

	// Routes for net segments
	api.GET("/segments", netSegHandler.GetAllNetSegments)
	api.POST("/segments", netSegHandler.CreateNetSegment)
	api.PUT("/segments/:id", netSegHandler.UpdateNetSegment)

	// Routes for hosts
	api.GET("/hosts", hostHandler.GetAllHosts)
	api.POST("/hosts", hostHandler.CreateHost)
	api.PUT("/hosts/:id", hostHandler.UpdateHost)
}
