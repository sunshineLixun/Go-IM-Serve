package main

import (
	"gin-chat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	
  // Migrate the schema
  db.AutoMigrate(&models.UserBasic{})

	user := models.UserBasic{
		Name: "1232",
		Password: "213123",
	}

  // Create
  db.Create(&user)


  // Update - update product's price to 200
  db.Model(&user).Update("Phone", "123123213212")

}