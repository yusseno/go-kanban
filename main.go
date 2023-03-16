package main

import (
	"go-kanban/app/handler"
	"go-kanban/app/repository"
	"go-kanban/app/service"
	"go-kanban/db"

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

	r := gin.Default()

	r.GET("login", apiHandler.UserAPIHandler.UserLogin)

	r.Run() // listen and serve on 0.0.0.0:8080
}
