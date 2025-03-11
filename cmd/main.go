package main

import (
	"github.com/labstack/echo/v4"

	"github.com/dr-aw/netseg-api/config"
	"github.com/dr-aw/netseg-api/internal/handler"
	"github.com/dr-aw/netseg-api/internal/repo"
	"github.com/dr-aw/netseg-api/internal/service"
	"github.com/dr-aw/netseg-api/internal/logger"
)

func main() {
	logger.InitLogger()
	log := logger.Logger
	cfg := config.LoadConfig()
	log.Info("Configuration loaded")
	db := repo.InitDB(cfg)
	log.Info("Database initialized")
	e := echo.New()
	log.Info("Echo server started")

	netSegBaseRepo := repo.NewNetSegmentRepo(db)
	netSegQueryRepo := repo.NewNetSegmentRepo(db)
	hostBaseRepo := repo.NewHostRepo(db)
	hostQueryRepo := repo.NewHostRepo(db)

	netSegService := service.NewNetSegmentService(netSegBaseRepo, netSegQueryRepo, hostQueryRepo)
	hostService := service.NewHostService(hostBaseRepo, hostQueryRepo, netSegQueryRepo)

	netSegHandler := handler.NewNetSegmentHandler(netSegService)
	hostHandler := handler.NewHostHandler(hostService)

	handler.RegisterRoutes(e, netSegHandler, hostHandler)

	// start the server
	e.Logger.Fatal(e.Start(":8080"))
}
