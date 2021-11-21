package http

import (
	"github.com/Revazashvili/easer/template"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, u template.UseCase) {
	h := NewHandler(u)
	temp := router.Group("template")
	{
		temp.GET("/", h.Get)
		temp.GET("/:id", h.GetById)
		temp.POST("/", h.Insert)
		temp.PUT("/:id", h.Update)
		temp.DELETE("/:id", h.Delete)
		temp.GET("/render/:id", h.Render)
	}
}
