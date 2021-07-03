package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() error {
	h := os.Getenv("MYSQL_HOST")
	u := os.Getenv("MYSQL_USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
	p := os.Getenv("MYSQL_PORT")
	d := os.Getenv("MYSQL_DATABASE")
	dsn := u + ":" + pwd + "@tcp(" + h + ":" + p + ")/" + d + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err

	}
	db = dbConnection
	return nil

}

func GetDB() *gorm.DB {
	return db
}
