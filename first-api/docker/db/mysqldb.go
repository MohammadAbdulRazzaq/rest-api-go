package db

import (
	"fmt"
	// config "learngo/restapiserver/pkg/movies/configs"
	config "first-api/config"
	"first-api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ReturnDB() (*gorm.DB, error) {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err.Error())
	}
	fmt.Println("______________________")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBLink, config.DBPort, config.DBName)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying *sql.DB: %w", err)
	}
	DB.AutoMigrate(&models.Mobile{})
	DB.AutoMigrate()

	fmt.Println("DB connected at:", sqlDB)
	return DB, nil
}
