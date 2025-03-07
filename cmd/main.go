package main

import (
	"github.com/dr-aw/netseg-api/internal/handler"
	"github.com/dr-aw/netseg-api/internal/repo"
	"github.com/dr-aw/netseg-api/internal/service"
)

func main() {
	db := repo.InitDB()
	_ = db

	var serv service.NetSegmentService
	_ = serv

	var handler handler.NetSegmentHandler
	_ = handler
}
