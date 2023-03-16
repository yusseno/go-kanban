package handler

import (
	"fmt"
	"go-kanban/app/user/entity"
	"go-kanban/app/user/service"
	"go-kanban/security"
	"go-kanban/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	UserLogin(c *gin.Context)
	UserRegister(c *gin.Context)
	UserLogout(c *gin.Context)
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

	resUser, err := u.userService.UserLogin(user)
	if err != nil {
		res := map[string]string{"message": err.Error()}
		c.JSON(200, res)
		return
	}
	// fmt.Println(user)
	getToken, err := session.SessionStore.Get(user.Email)
	if err != nil {
		tokenString, err := security.GenerateJWT(resUser.Email, resUser.Fullname, resUser.Is_Role)
		if err != nil {
			res := map[string]string{"message": "failed to generate token"}
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		session.SessionStore.Set(user.Email, session.SessionToken{
			TokenString: tokenString,
		})

		res := map[string]string{"email": user.Email, "token": tokenString, "message": "You are logged in"}
		c.JSON(http.StatusOK, res)
		return
	} else {
		res := map[string]string{"email": user.Email, "token": getToken.TokenString, "message": "You are logged in"}
		c.JSON(http.StatusOK, res)
		return
	}
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

func (u *userAPI) UserLogout(c *gin.Context) {
	key := c.MustGet("objek").(string)
	fmt.Println("ini dari middelware", key)
	err := session.SessionStore.Del(key)
	if err != nil {
		res := map[string]string{"message": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
}
