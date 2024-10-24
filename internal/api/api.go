package api

import (
	"github.com/edimarlnx/html2pdf-service/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var (
	authCode string = "dev"
	authUser string = "dev"
)

func init() {
	authUser = utils.GetEnv("AUTH_USER", "dev")
	authCode = utils.GetEnv("AUTH_CODE", "dev")
}

func BasicAuth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == authUser && password == authCode {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, &Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Info: map[string]string{
			"error": "Namespace not found",
		},
	})
}

func Auth(c *gin.Context) {
	println(c.FullPath())
	if strings.HasPrefix(c.FullPath(), "/_health") {
		c.Next()
		return
	}
	apiKey := c.GetHeader("x-api-key")
	if apiKey == "" {
		apiKey, _ = c.GetQuery("api-key")
	}
	if apiKey != authCode {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &Response{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
			Info: map[string]string{
				"error": "Unauthorized",
			},
		})
		return
	}
	c.Next()
}

func Start() {
	engine := gin.Default()
	routerGroup := engine.Group("/")
	//routerGroup.Use(BasicAuth)
	routerGroup.Use(Auth)
	wsHealth(routerGroup)
	fromUrl(routerGroup)
	fromHtmlContent(routerGroup)

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Error starting server %v", err)
		return
	}
}
