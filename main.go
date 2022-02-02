package main

import (
	"github.com/manishdangi98/banking/app"
	"github.com/manishdangi98/banking/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()
}
