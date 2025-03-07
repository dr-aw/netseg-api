package main

import (
	"github.com/labstack/echo/v4"

	"github.com/dr-aw/netseg-api/config"
	"github.com/dr-aw/netseg-api/internal/handler"
	"github.com/dr-aw/netseg-api/internal/repo"
	"github.com/dr-aw/netseg-api/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	db := repo.InitDB(cfg)
	e := echo.New()

	netSegRepo := repo.NewNetSegmentRepo(db)
	hostRepo := repo.NewHostRepo(db)

	netSegService := service.NewNetSegmentService(netSegRepo)
	hostService := service.NewHostService(hostRepo, netSegService)

	netSegHandler := handler.NewNetSegmentHandler(netSegService)
	hostHandler := handler.NewHostHandler(hostService)

	handler.RegisterNetSegmentRoutes(e, netSegHandler)
	handler.RegisterHostRoutes(e, hostHandler)

	// start the server
	e.Logger.Fatal(e.Start(":8080"))
}
