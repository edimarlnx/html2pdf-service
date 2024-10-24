package api

import (
	"fmt"
	"github.com/edimarlnx/html2pdf-service/internal/render"
	"github.com/gin-gonic/gin"
	"net/http"
)

func fromUrl(e *gin.RouterGroup) {
	e.GET("/create-pdf", func(c *gin.Context) {
		var request Request
		bindError := c.BindQuery(&request)
		if bindError != nil {
			errMsg := fmt.Sprintf("Error on process request. %s", bindError)
			fmt.Println(errMsg)
			badResponse(c, errMsg)
			return
		}
		if request.DownloadFileName == "" {
			request.DownloadFileName = "generated-file.pdf"
		}
		processPDFRequest(c, request)
	})
}

func fromHtmlContent(e *gin.RouterGroup) {
	e.POST("/create-pdf", func(c *gin.Context) {
		var request Request
		bindError := c.BindQuery(&request)
		if bindError != nil {
			errMsg := fmt.Sprintf("Error on process request. %s", bindError)
			fmt.Println(errMsg)
			badResponse(c, errMsg)
			return
		}
		request.Content, _ = c.GetRawData()
		if request.DownloadFileName == "" {
			request.DownloadFileName = "generated-file.pdf"
		}

		processPDFRequest(c, request)
	})
}

func processPDFRequest(c *gin.Context, request Request) {
	var pdfData []byte
	var err error
	if request.Content != nil {
		pdfData, err = render.PDFFromContent(request.Content, request.WaitForSelector)
	} else {
		pdfData, err = render.PDF(request.Url, request.Headers, request.WaitForSelector)
	}
	if err != nil {
		errMsg := fmt.Sprintf("Error on create PDF. %s", err.Error())
		fmt.Println(errMsg)
		badResponse(c, errMsg)
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", request.DownloadFileName))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(pdfData)))
	c.Data(http.StatusOK, "application/octet-stream", pdfData)
}
