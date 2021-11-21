package http

import (
	"github.com/Revazashvili/easer/pdf"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	useCase pdf.UseCase
}

func NewHandler(useCase pdf.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Render(c *gin.Context) {
	id := c.Param("id")
	var data interface{}
	err := c.BindJSON(&data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	bytes, ok := h.useCase.Render(id, data)
	if !ok {
		c.JSON(http.StatusBadRequest, err)
	}
	c.Writer.Header().Set("Content-Type", "application/pdf")
	c.JSON(http.StatusOK, bytes)
}
