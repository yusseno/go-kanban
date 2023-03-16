package handler

import (
	"go-kanban/app/user/entity"
	"go-kanban/app/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	UserLogin(c *gin.Context)
	UserRegister(c *gin.Context)
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
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		res := map[string]string{"message": "json not match"}
		c.JSON(200, res)
		return
	}

	if user.Fullname == "" && user.Password == "" && user.Email == "" {
		res := map[string]string{"message": "input is empty"}
		c.JSON(200, res)
		return
	}

	if err := u.userService.UserLogin(user); err != nil {
		res := map[string]string{"message": err.Error()}
		c.JSON(200, res)
		return
	}

	res := entity.ResLogin{
		Email:   user.Email,
		Message: "Login berhasil",
	}
	c.JSON(http.StatusCreated, res)
}

func (u *userAPI) UserRegister(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		res := map[string]string{"message": "json not match"}
		c.JSON(200, res)
		return
	}

	if user.Fullname == "" && user.Password == "" && user.Email == "" {
		res := map[string]string{"message": "input is empty"}
		c.JSON(200, res)
		return
	}

	if err := u.userService.UserRegister(user); err != nil {
		res := map[string]string{"message": err.Error()}
		c.JSON(200, res)
		return
	}

	res := entity.ResLogin{
		Email:   user.Email,
		Message: "Register berhasil",
	}
	c.JSON(http.StatusCreated, res)
}
