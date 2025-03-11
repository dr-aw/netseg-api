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

	netSegBaseRepo := repo.NewNetSegmentRepo(db)
	netSegQueryRepo := repo.NewNetSegmentRepo(db)
	hostBaseRepo := repo.NewHostRepo(db)
	hostQueryRepo := repo.NewHostRepo(db)

	// ✅ Создаём `NetSegmentService` без `HostService`
	netSegService := service.NewNetSegmentService(netSegBaseRepo, netSegQueryRepo, hostQueryRepo)

	// ✅ `HostService` остаётся без изменений
	hostService := service.NewHostService(hostBaseRepo, hostQueryRepo, netSegQueryRepo)

	netSegHandler := handler.NewNetSegmentHandler(netSegService)
	hostHandler := handler.NewHostHandler(hostService)

	handler.RegisterRoutes(e, netSegHandler, hostHandler)

	// start the server
	e.Logger.Fatal(e.Start(":8080"))
}
