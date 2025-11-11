package main

import (
	"github.com/rodrigofnobrega/gopportunities/config"
	"github.com/rodrigofnobrega/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errf("config initialization erro: %v", err)
		return
	}
	router.Initialize()
}
