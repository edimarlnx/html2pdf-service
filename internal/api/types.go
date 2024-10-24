package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message,omitempty"`
	Info    map[string]string `json:"info,omitempty"`
}

type Request struct {
	Url              string                 `form:"url,omitempty"`
	Headers          map[string]interface{} `form:"headers"`
	DownloadFileName string                 `form:"downloadFileName,omitempty"`
	//SendToURL        string                 `form:"sendToURL,omitempty"`
	WaitForSelector string `form:"waitForSelector,omitempty"`
	Content         []byte
}

func badResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, &Response{
		Status:  http.StatusBadRequest,
		Message: message,
	})
}
