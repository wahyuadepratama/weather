package main

import (
	"github.com/wahyuadepratama/weather/src/config"

	"github.com/wahyuadepratama/weather/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	//run all routes
	routes.Routes()
}
