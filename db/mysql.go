package db

import (
	"fmt"
	entityKanban "go-kanban/app/kanban/entity"
	entityUser "go-kanban/app/user/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "admin"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_NAME     = "go-kanban"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	db = ConnectDB()
	return db
}
func ConnectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return nil
	}
	// fmt.Println("Connected to database")
	db.AutoMigrate(&entityUser.User{}, &entityKanban.Category{})
	db.Delete(&entityUser.User{}, &entityKanban.Category{})
	return db
}
