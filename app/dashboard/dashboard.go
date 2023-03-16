package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	res := map[string]string{"message": "ada sebagai admin"}
	c.JSON(http.StatusOK, res)
}
