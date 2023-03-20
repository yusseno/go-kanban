package handler

import (
	"go-kanban/app/kanban/service"

	"github.com/gin-gonic/gin"
)

type CategoryAPI interface {
	CreateCategory(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{
		categoryService: categoryService,
	}
}

func (api *categoryAPI) CreateCategory(c *gin.Context) {
	c.JSON(200, "success")
}
