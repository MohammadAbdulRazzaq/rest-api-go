package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "password",
		DBName:   "mydb",
	}
	return &dbConfig
}

func InitDB() {
	dbConfig := BuildDBConfig()
	dbURL := DbURL(dbConfig)
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB = db
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
