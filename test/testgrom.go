package main

import (
	"fmt"
	"gin_Chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:12345678@tcp(localhost:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "张三"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1))

	// Update - update product's price to 200
	db.Model(user).Update("Password", "1234")

}
