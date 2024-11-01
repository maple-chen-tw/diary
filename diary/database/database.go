package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "maple"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3307
	DATABASE = "diary_app"
)

func InitializeDB() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	return db
}
