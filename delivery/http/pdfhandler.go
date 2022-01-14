package delivery

import (
	"github.com/Revazashvili/easer/renderers"
	"github.com/Revazashvili/easer/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterPdfHTTPEndpoints(router *gin.RouterGroup, s storage.Storage, r renderers.Renderer) {
	h := newHandler(s, r)
	temp := router.Group("pdf")
	{
		temp.GET("/render/:id", h.Render)
	}
}

type pdfHandler struct {
	storage  storage.Storage
	renderer renderers.Renderer
}

func newHandler(s storage.Storage, r renderers.Renderer) *pdfHandler {
	return &pdfHandler{s, r}
}

func (h *pdfHandler) Render(c *gin.Context) {
	id := c.Param("id")
	var data interface{}
	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	template, err := h.storage.Get(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	bytes, ok := h.renderer.Render(template, data)
	if !ok {
		c.JSON(http.StatusBadRequest, err)
	}
	c.Writer.Header().Set("Content-Type", "application/pdf")
	c.JSON(http.StatusOK, bytes)
}
