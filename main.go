package main

import (
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/config"
	"github.com/mufasa-dev/Wallet-flow-api-in-Golang/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error %v", err)
		return
	}

	router.Initialize()
}
