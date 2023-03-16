package handler

import (
	"go-kanban/app/service"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	UserLogin(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{
		userService: userService,
	}
}

func (u *userAPI) UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
