package http

import (
	"github.com/Revazashvili/easer/template"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase template.UseCase
}

func NewHandler(useCase template.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type getResponse struct {
	Templates []*Template `json:"templates"`
}

func(h * Handler) Get(c *gin.Context)  {
	ts,err := h.useCase.All()
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,&getResponse{
		Templates: asResponses(ts),
	})
}

type getByIdResponse struct {
	Template *Template `json:"template"`
}

func(h *Handler) GetById(c *gin.Context)  {
	id := c.Param("id")
	if id == "" || len(id) <= 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	t,err := h.useCase.Find(id)
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,&getByIdResponse{
		Template: asResponse(t),
	})
}

type insertResponse struct {
	Id string `json:"id"`
}

func (h *Handler) Insert(c *gin.Context)  {
	t := new(Template)
	err := c.BindJSON(&t)
	if err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id,err := h.useCase.Insert(asDomain(t))
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,&insertResponse{
		Id: id,
	})
}

type updateResponse struct {
	Id string `json:"id"`
}

func (h *Handler) Update(c *gin.Context)  {
	t := new(Template)
	err := c.BindJSON(&t)
	if err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id,err := h.useCase.Update(asDomain(t))
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,&updateResponse{
		Id: id,
	})
}

func (h *Handler) Delete(g *gin.Context) {
	err := h.useCase.Delete(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.AbortWithStatus(http.StatusOK)
}
