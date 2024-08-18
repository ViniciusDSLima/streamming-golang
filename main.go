package main

import (
	"github.com/ViniciusDSLima/streamming-golang/config"
	"github.com/ViniciusDSLima/streamming-golang/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errf("Config initialization failed: %v", err)
		return
	}

	router.Initialize()
}
