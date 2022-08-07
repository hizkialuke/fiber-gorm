package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	DBConn = db
	// 	sqlDB, err := DBConn.DB()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer sqlDB.Close()
}
