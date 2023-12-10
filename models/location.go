package models

type Location struct {
	Id        uint   `json:"id" gorm:"primary_key;autoIncrement:true"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Category  string `json:"category"`
}

type LocationInput struct {
	Name      string `json:"name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Latitude  string `json:"latitude" binding:"required"`
	Longitude string `json:"longitude" binding:"required"`
	Category  string `json:"category" binding:"required"`
}
