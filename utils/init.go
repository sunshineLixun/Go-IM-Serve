package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
}

func InitMySQL() {
	_db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = _db
}

func GetDb() *gorm.DB {
	return db
}
