package http

import (
	"github.com/Revazashvili/easer/pdf"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, pdf pdf.UseCase) {
	h := NewHandler(pdf)
	temp := router.Group("pdf")
	{
		temp.GET("/render/:id", h.Render)
	}
}
