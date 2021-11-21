package http

import (
	"github.com/Revazashvili/easer/pdf"
	"github.com/gin-gonic/gin"
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
	bytes, err := h.useCase.Render(id, data)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
	}
	c.Writer.Header().Set("Content-Type", "application/pdf")
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, bytes)
}
