package main

import (
	"github.com/manishdangi98/banking-lib/logger"
	"github.com/manishdangi98/banking/app"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()
}
