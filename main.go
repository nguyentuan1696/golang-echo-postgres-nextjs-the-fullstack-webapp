package main

import (
	"go-api-starter/core/logger"
	"go-api-starter/core/server"
)

func main() {
	if err := server.Run(); err != nil {
		logger.Error("run server error", err)
	}
}
