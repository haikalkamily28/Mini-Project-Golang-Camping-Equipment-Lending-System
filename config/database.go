package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type ConfigDB struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

var DB *gorm.DB

// ConnectDb initializes a connection to the database with connection pooling and timeout settings
func ConnectDb() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read database configurations from environment variables
	configDB := ConfigDB{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
	}

	// Set up the connection string with timeout settings
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s", 
		configDB.User, 
		configDB.Password, 
		configDB.Host, 
		configDB.Port, 
		configDB.Name)

	// Open the database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Set up connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get SQL DB instance: %v", err)
	}

	// Set the maximum number of open connections and idle connections
	sqlDB.SetMaxOpenConns(10)  // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(1 * time.Hour) // Set connection max lifetime

	// Set the global DB instance for use throughout the application
	DB = db
	log.Println("Database connected successfully")
}
