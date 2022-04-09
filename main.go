package main

import (
	"github.com/crobatair/banking/app"
	"github.com/crobatair/banking/logger"
)

func main() {

	logger.Info("Starting banking service...")
	app.StartApp()

}
