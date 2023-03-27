package models

import (
	"time"

	"gorm.io/gorm"

	"gin-chat/utils"
)

type UserBasic struct {
	gorm.Model
	Name         string
	Password     string
	Phone        string
	Email        string
	ClientIp     string
	ClientPort   string
	LoginTime    time.Time
	LoginOutTime time.Time
}

// GetUserList 获取所有用户
func GetUserList() []*UserBasic {
	users := make([]*UserBasic, 10)
	utils.GetDb().Find(&users)
	return users
}
