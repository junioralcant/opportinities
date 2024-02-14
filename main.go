package main

import (
	"github.com/junioralcant/opportinities.git/config"
	"github.com/junioralcant/opportinities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
