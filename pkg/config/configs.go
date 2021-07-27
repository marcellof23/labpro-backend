package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	h := os.Getenv("MYSQL_HOST")
	u := os.Getenv("MYSQL_USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
	p := os.Getenv("MYSQL_PORT")
	d := os.Getenv("MYSQL_DATABASE")
	log.Println(h, u, pwd, p, d)
	dsn := u + ":" + pwd + "@tcp(" + h + ":" + p + ")/" + d + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println(dsn)
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err

	}
	DB = dbConnection
	return nil
}

var (
	DefaultPort = "8080"
)
