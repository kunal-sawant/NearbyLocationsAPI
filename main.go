package main

import (
	"example/NearbyLocationsAPI/controller"
	"example/NearbyLocationsAPI/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initializers.ConnectDB()

	r.GET("/locations", controller.FindLocation)
	r.GET("/locations/:category", controller.GetLocByCategory)
	r.POST("/locations", controller.CreateLocation)

	r.Run("localhost:8080")
}
