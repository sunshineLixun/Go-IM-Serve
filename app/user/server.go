package user

import (
	"gin-chat/models"
	"github.com/gin-gonic/gin"
)

func QueryHandler(context *gin.Context) {

	users := models.GetUserList()
	context.JSON(200, gin.H{
		"list":    users,
		"message": "GET",
	})
}
