package config

import (
	"gorm.io/gorm"
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
		DBName:   "MYDB",
	}
	return &dbConfig
}

// func InitDB() (*gorm.DB, error) {
// 	dbConfig := BuildDBConfig()
// 	dbURL := DbURL(dbConfig)
// 	fmt.Println("Database URL:", dbURL) // Print the database URL for debugging

// 	db, err := gorm.Open("mysql", dbURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to database: %w", err)
// 	}
// 	return db, nil
// }

// func DbURL(dbConfig *DBConfig) string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }
