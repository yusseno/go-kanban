package main

import (
	"go-kanban/app/dashboard"
	"go-kanban/app/home"
	"go-kanban/app/middelware"
	"go-kanban/app/user/handler"
	"go-kanban/app/user/repository"
	"go-kanban/app/user/service"
	"go-kanban/db"
	"go-kanban/session"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	UserAPIHandler handler.UserAPI
}

func main() {
	//repository
	userRepo := repository.NewUserRepository(db.ConnectDB())

	//service
	userService := service.NewUserService(userRepo)

	//handlers
	userAPIHandler := handler.NewUserAPI(userService)
	apiHandler := APIHandler{
		UserAPIHandler: userAPIHandler,
	}

	session.SessionStore = session.NewRedisStore()

	r := gin.Default()

	// group client user
	r.GET("login", apiHandler.UserAPIHandler.UserLogin)
	r.POST("register", apiHandler.UserAPIHandler.UserRegister)
	r.GET("logout", middelware.AuthzUser(), apiHandler.UserAPIHandler.UserLogout)
	r.GET("user", middelware.AuthzUser(), home.Home)
	r.GET("dashboard", middelware.AuthzAdmin(), dashboard.Dashboard)

	r.Run() // listen and serve on 0.0.0.0:8080
}
