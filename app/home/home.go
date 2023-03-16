package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	res := map[string]string{"message": "ada sebagai user"}
	c.JSON(http.StatusOK, res)
}
