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
	netSegService := service.NewNetSegmentService(netSegRepo)
	netSegHandler := handler.NewNetSegmentHandler(netSegService)

	handler.RegisterNetSegmentRoutes(e, netSegHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
