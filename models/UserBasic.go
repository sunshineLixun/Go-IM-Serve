package models

import (
	"gorm.io/gorm"
)


type UserBasic struct {
	gorm.Model
	Name string
	Password string
	Phone string
	Email string
	ClientIp string
	ClientPort string
	LoginTime uint64
	LogoutTime uint64
	
}