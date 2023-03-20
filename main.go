package main

import (
	"go-kanban/app/dashboard"
	"go-kanban/app/home"
	kanbanHandler "go-kanban/app/kanban/handler"
	kanbanRepo "go-kanban/app/kanban/repository"
	kanbanService "go-kanban/app/kanban/service"
	"go-kanban/app/middelware"
	userHandler "go-kanban/app/user/handler"
	userRepo "go-kanban/app/user/repository"
	userService "go-kanban/app/user/service"
	"go-kanban/db"
	"go-kanban/session"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	UserAPIHandler     userHandler.UserAPI
	CategoryAPIHandler kanbanHandler.CategoryAPI
}

func main() {
	//repository
	userRepo := userRepo.NewUserRepository(db.ConnectDB())
	categoryRepo := kanbanRepo.NewCategoryRepository(db.ConnectDB())

	//service
	userService := userService.NewUserService(userRepo)
	categoryService := kanbanService.NewCategoryService(categoryRepo)

	//handlers
	userAPIHandler := userHandler.NewUserAPI(userService)
	categoryAPIHandler := kanbanHandler.NewCategoryAPI(categoryService)
	apiHandler := APIHandler{
		UserAPIHandler:     userAPIHandler,
		CategoryAPIHandler: categoryAPIHandler,
	}

	session.SessionStore = session.NewRedisStore()

	r := gin.Default()

	// group client user
	r.GET("login", apiHandler.UserAPIHandler.UserLogin)
	r.POST("register", apiHandler.UserAPIHandler.UserRegister)
	r.GET("logout", middelware.AuthzUser(), apiHandler.UserAPIHandler.UserLogout)

	r.GET("user", middelware.AuthzUser(), home.Home)
	r.GET("dashboard", middelware.AuthzAdmin(), dashboard.Dashboard)

	r.POST("category", apiHandler.CategoryAPIHandler.CreateCategory)

	r.Run() // listen and serve on 0.0.0.0:8080
}
