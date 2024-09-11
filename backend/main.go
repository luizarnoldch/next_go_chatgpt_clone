package main

import (
	"main/app"

	fiberlog "github.com/gofiber/fiber/v2/log"
)

func main() {
	fiberlog.Info("Server Starting")
	app.Start()
	fiberlog.Info("Server Finished")	
}