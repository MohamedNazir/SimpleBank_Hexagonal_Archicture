package main

import (
	"github.com/MohamedNazir/SimpleBank/api"
	"github.com/MohamedNazir/SimpleBank/logger"
)

func main() {
	logger.Info("Starting Application")
	api.StartApplication()

}
