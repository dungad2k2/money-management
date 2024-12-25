package models
import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
)
var DB *gorm.DB

func ConnectDB() {
	dbName := "User"
	
	noDBDsn := "buidung:1@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	noDBConn, err := gorm.Open(mysql.Open(noDBDsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL server!")
	}else {
		fmt.Println("Connected to MySQL server!")
	}

	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := noDBConn.Exec(createDBSQL).Error; err != nil {
		panic(fmt.Sprintf("Failed to create database %s: %v", dbName, err))
	}

	dsn := fmt.Sprintf("buidung:1@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{}, &Expense{})

	DB = database
}

