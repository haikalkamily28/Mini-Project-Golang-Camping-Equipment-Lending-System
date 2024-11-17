package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
	"log"
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

func ConnectDb() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    configDB := ConfigDB{
        Host:     os.Getenv("DATABASE_HOST"),
        User:     os.Getenv("DATABASE_USER"),
        Password: os.Getenv("DATABASE_PASSWORD"),
        Port:     os.Getenv("DATABASE_PORT"),
        Name:     os.Getenv("DATABASE_NAME"),
    }

    dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
        configDB.User,
        configDB.Password,
        configDB.Host,
        configDB.Port,
        configDB.Name)

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
}
