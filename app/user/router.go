package user

import (
	"github.com/gin-gonic/gin"
)

// Routers
// @Schemes https
// @Description 用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {json} json{"code", "message"}
// @Router /user/getUsers [get]
func Routers(e *gin.Engine) {
	e.GET("/user/getUsers", QueryHandler)
}
