package database

import (
	"fmt"
	"log"
	"time"
	"learn-microservices/user-service/config"
	"learn-microservices/user-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// Load konfigurasi database dari .env
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "")

	// Format DSN untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Koneksi ke database
	var err error
	for i := 1; i <= 10; i++ {

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, _ := DB.DB()
			err = sqlDB.Ping()
		}

		if err == nil {
			fmt.Println("Database connected successfully!")
			break
		}

		fmt.Printf("Database not ready, retrying... (%d/10)\n", i)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// **Auto Migrate Models**
	err = DB.AutoMigrate(
		&models.User{},
		// &models.Course{},
		// &models.UserCourse{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated successfully!")
}
