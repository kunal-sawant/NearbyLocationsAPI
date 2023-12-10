package controller

import (
	"example/NearbyLocationsAPI/initializers"
	"example/NearbyLocationsAPI/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindLocation(c *gin.Context) {
	now := time.Now().UnixNano()
	var locations []models.Location
	initializers.DB.Find(&locations)
	after := time.Now().UnixNano()
	c.JSON(http.StatusOK, gin.H{"time_ns": after - now, "data": locations})
}

func CreateLocation(c *gin.Context) {
	now := time.Now().UnixNano()
	var inputLoc models.LocationInput
	if err := c.ShouldBindJSON(&inputLoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location := models.Location{
		Name:      inputLoc.Name,
		Address:   inputLoc.Address,
		Latitude:  inputLoc.Latitude,
		Longitude: inputLoc.Longitude,
		Category:  inputLoc.Category,
	}

	// Migrate the schema
	err := initializers.DB.AutoMigrate(&models.Location{})
	if err != nil {
		fmt.Println(err)

	}

	initializers.DB.Create(&location)
	after := time.Now().UnixNano()

	c.JSON(http.StatusOK, gin.H{"time_ns": after - now, "data": location})
}

func GetLocByCategory(c *gin.Context) {
	now := time.Now().UnixNano()
	category := c.Param("category")
	fmt.Println("category:" + category)

	var locations []models.Location
	initializers.DB.Where(&models.Location{Category: category}).Find(&locations)
	after := time.Now().UnixNano()
	c.JSON(http.StatusOK, gin.H{"time_ns": after - now, "locations": locations})
}
