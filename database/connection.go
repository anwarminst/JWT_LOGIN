package database

import (
	"github.com/anwarminst/jwt-auth-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to MySQL database
func ConnectionDB() (*gorm.DB, error) {

	// Database configuration
	dsn := "root:ZXcvbnm@123@tcp(localhost:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	DB = db
	// Automatically migrate the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
