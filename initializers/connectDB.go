package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func ConnectDB(config *Config)
func ConnectDB() {
	var err error
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	dsn := "host=localhost user=postgres password=Zoro@1.11B dbname=AssessmentDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
		fmt.Println("Failed to connect to the Database")
		return
	}
	fmt.Println("Connected Successfully to the Database")
	// DB.Find(nil)
}
