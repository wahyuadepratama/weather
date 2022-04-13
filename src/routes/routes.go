package routes

import (
	"github.com/wahyuadepratama/weather/src/controllers"

	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func Routes() {

	route := gin.Default()

	route.POST("/user/login", controllers.UserLogin)
	route.POST("/user/register", controllers.UserRegister)

	route.GET("/weather/all", controllers.ShowAllWeather)
	route.GET("/weather/update", controllers.UpdateWeather)

	// Run route whenever triggered
	route.Run()
}
