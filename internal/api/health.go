package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func wsHealth(e *gin.RouterGroup) {
	e.GET("/_health", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, &Response{
			Status:  http.StatusOK,
			Message: "up",
		})
	})
}
