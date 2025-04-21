package main

import (
	"thichlab-backend-slowpoke/core/logger"
	"thichlab-backend-slowpoke/core/server"
)

func main() {
	if err := server.Run(); err != nil {
		logger.Error("run server error", err)
	}
}
