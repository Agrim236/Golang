package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
	"golang-notes-api/models"
)

var DB *gorm.DB

func ConnectDB() error {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Println("❌ Error loading .env file")
		return err
	}
	
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	
	// Validate environment variables
	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		log.Println("❌ Database environment variables not properly set")
		log.Printf("DB_USER: %s, DB_PASSWORD: [hidden], DB_HOST: %s, DB_PORT: %s, DB_NAME: %s\n", 
			user, host, port, dbname)
		return fmt.Errorf("database environment variables not properly set")
	}
	
	// Try different connection methods
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true&multiStatements=true",
		user, password, host, port, dbname,
	)
		log.Printf("🔄 Attempting to connect to MySQL with host: %s, port: %s, user: %s, database: %s",
		host, port, user, dbname)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Database connection failed: %v\nDSN: %s\n", err, dsn)
		return err
	}
	
	DB = db
	log.Println("✅ Database connection established.")
	
	// Auto-migrate the models
	log.Println("🔄 Running database migrations...")
	err = DB.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Printf("❌ Database migration failed: %v", err)
		return err
	}
	
	log.Println("✅ Database migrations completed successfully.")
	return nil
}