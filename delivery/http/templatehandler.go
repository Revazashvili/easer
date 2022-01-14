package delivery

import (
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/parsers"
	"github.com/Revazashvili/easer/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterTemplateHTTPEndpoints(router *gin.RouterGroup, s storage.Storage, p parsers.Parser) {
	h := newTemplateHandler(s, p)
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

type templateHandler struct {
	storage storage.Storage
	parser  parsers.Parser
}

func newTemplateHandler(s storage.Storage, p parsers.Parser) *templateHandler {
	return &templateHandler{s, p}
}

type getResponse struct {
	Templates []template `json:"templates"`
}

func (h *templateHandler) Get(c *gin.Context) {
	ts, err := h.storage.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, &getResponse{
		Templates: asResponses(ts),
	})
}

type getByIdResponse struct {
	Template template `json:"mongo"`
}

func (h *templateHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" || len(id) <= 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	t, err := h.storage.Get(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, &getByIdResponse{
		Template: asResponse(t),
	})
}

type insertResponse struct {
	Id string `json:"id"`
}

func (h *templateHandler) Insert(c *gin.Context) {
	var t template
	err := c.BindJSON(&t)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := h.storage.Add(asDomain(t))
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, &insertResponse{
		Id: id,
	})
}

type updateResponse struct {
	Status bool `json:"status,omitempty"`
}

func (h *templateHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" || len(id) <= 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var t template
	err := c.BindJSON(&t)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	status, err := h.storage.Update(id, asDomain(t))
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, &updateResponse{
		Status: status,
	})
}

func (h *templateHandler) Delete(c *gin.Context) {
	err := h.storage.Delete(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
	}
	c.AbortWithStatus(http.StatusOK)
}

type renderResponse struct {
	Html string
}

func (h *templateHandler) Render(c *gin.Context) {
	id := c.Param("id")
	var data interface{}
	err := c.BindJSON(&data)
	temp, err := h.storage.Get(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err.Error())
	}
	html, ok := h.parser.Parse(temp.Name, temp.TemplateBody, data)
	if !ok {
		c.String(http.StatusInternalServerError, "%s", "Can't parse html")
	}
	c.JSON(http.StatusOK, &renderResponse{
		Html: html,
	})
}

type template struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Owner        string  `json:"owner"`
	TemplateBody string  `json:"template_body"`
	Options      options `json:"options"`
}

type headerAndFooterOptions struct {
	FooterCenter   string  `json:"footer_center"`
	FooterFontName string  `json:"footer_font_name"`
	FooterFontSize uint    `json:"footer_font_size"`
	FooterHTML     string  `json:"footer_html"`
	FooterLeft     string  `json:"footer_left"`
	FooterLine     bool    `json:"footer_line"`
	FooterRight    string  `json:"footer_right"`
	FooterSpacing  float64 `json:"footer_spacing"`
	HeaderCenter   string  `json:"header_center"`
	HeaderFontName string  `json:"header_font_name"`
	HeaderFontSize uint    `json:"header_font_size"`
	HeaderHTML     string  `json:"header_html"`
	HeaderLeft     string  `json:"header_left"`
	HeaderLine     bool    `json:"header_line"`
	HeaderRight    string  `json:"header_right"`
	HeaderSpacing  float64 `json:"header_spacing"`
}

type margin struct {
	Top    uint `json:"top"`
	Bottom uint `json:"bottom"`
	Left   uint `json:"left"`
	Right  uint `json:"right"`
}

type options struct {
	Grayscale            bool                   `json:"grayscale"`
	Dpi                  uint                   `json:"dpi"`
	DisplayHeaderFooter  bool                   `json:"display_header_footer"`
	HeaderFooterOptions  headerAndFooterOptions `json:"header_footer_options"`
	PrintBackground      bool                   `json:"print_background"`
	Orientation          string                 `json:"orientation"`
	Format               string                 `json:"format"`
	EnableForms          bool                   `json:"enable_forms"`
	DisableExternalLinks bool                   `json:"disable_external_links"`
	DisableInternalLinks bool                   `json:"disable_internal_links"`
	NoBackground         bool                   `json:"no_background"`
	NoImages             bool                   `json:"no_images"`
	Margin               margin                 `json:"margin"`
}

func asResponse(t models.Template) template {
	return template{
		Id:           t.Id,
		Name:         t.Name,
		Description:  t.Description,
		Owner:        t.Owner,
		TemplateBody: t.TemplateBody,
		Options: options{
			DisableExternalLinks: t.Options.DisableExternalLinks,
			DisableInternalLinks: t.Options.DisableInternalLinks,
			Dpi:                  t.Options.Dpi,
			DisplayHeaderFooter:  t.Options.DisplayHeaderFooter,
			EnableForms:          t.Options.EnableForms,
			Format:               t.Options.Format,
			Grayscale:            t.Options.Grayscale,
			NoImages:             t.Options.NoImages,
			NoBackground:         t.Options.NoBackground,
			Orientation:          t.Options.Orientation,
			PrintBackground:      t.Options.PrintBackground,
			Margin: margin{
				Top:    t.Options.Margin.Top,
				Bottom: t.Options.Margin.Bottom,
				Left:   t.Options.Margin.Left,
				Right:  t.Options.Margin.Left,
			},
			HeaderFooterOptions: headerAndFooterOptions{
				FooterCenter:   t.Options.HeaderFooterOptions.FooterCenter,
				FooterHTML:     t.Options.HeaderFooterOptions.FooterHTML,
				FooterLeft:     t.Options.HeaderFooterOptions.FooterLeft,
				FooterLine:     t.Options.HeaderFooterOptions.FooterLine,
				FooterFontName: t.Options.HeaderFooterOptions.FooterFontName,
				FooterFontSize: t.Options.HeaderFooterOptions.FooterFontSize,
				FooterRight:    t.Options.HeaderFooterOptions.FooterRight,
				FooterSpacing:  t.Options.HeaderFooterOptions.FooterSpacing,
				HeaderCenter:   t.Options.HeaderFooterOptions.HeaderCenter,
				HeaderHTML:     t.Options.HeaderFooterOptions.HeaderHTML,
				HeaderLeft:     t.Options.HeaderFooterOptions.HeaderLeft,
				HeaderLine:     t.Options.HeaderFooterOptions.HeaderLine,
				HeaderFontName: t.Options.HeaderFooterOptions.HeaderFontName,
				HeaderFontSize: t.Options.HeaderFooterOptions.HeaderFontSize,
				HeaderRight:    t.Options.HeaderFooterOptions.HeaderRight,
				HeaderSpacing:  t.Options.HeaderFooterOptions.HeaderSpacing,
			},
		},
	}
}

func asResponses(ts []models.Template) []template {
	out := make([]template, len(ts))
	for i, t := range ts {
		out[i] = asResponse(t)
	}
	return out
}

func asDomain(t template) models.Template {
	return models.Template{
		Id:           t.Id,
		Name:         t.Name,
		Description:  t.Description,
		Owner:        t.Owner,
		TemplateBody: t.TemplateBody,
		Options: models.Options{
			DisableExternalLinks: t.Options.DisableExternalLinks,
			DisableInternalLinks: t.Options.DisableInternalLinks,
			Dpi:                  t.Options.Dpi,
			DisplayHeaderFooter:  t.Options.DisplayHeaderFooter,
			EnableForms:          t.Options.EnableForms,
			Format:               t.Options.Format,
			Grayscale:            t.Options.Grayscale,
			NoImages:             t.Options.NoImages,
			NoBackground:         t.Options.NoBackground,
			Orientation:          t.Options.Orientation,
			PrintBackground:      t.Options.PrintBackground,
			Margin: models.Margin{
				Top:    t.Options.Margin.Top,
				Bottom: t.Options.Margin.Bottom,
				Left:   t.Options.Margin.Left,
				Right:  t.Options.Margin.Left,
			},
			HeaderFooterOptions: models.HeaderAndFooterOptions{
				FooterCenter:   t.Options.HeaderFooterOptions.FooterCenter,
				FooterHTML:     t.Options.HeaderFooterOptions.FooterHTML,
				FooterLeft:     t.Options.HeaderFooterOptions.FooterLeft,
				FooterLine:     t.Options.HeaderFooterOptions.FooterLine,
				FooterFontName: t.Options.HeaderFooterOptions.FooterFontName,
				FooterFontSize: t.Options.HeaderFooterOptions.FooterFontSize,
				FooterRight:    t.Options.HeaderFooterOptions.FooterRight,
				FooterSpacing:  t.Options.HeaderFooterOptions.FooterSpacing,
				HeaderCenter:   t.Options.HeaderFooterOptions.HeaderCenter,
				HeaderHTML:     t.Options.HeaderFooterOptions.HeaderHTML,
				HeaderLeft:     t.Options.HeaderFooterOptions.HeaderLeft,
				HeaderLine:     t.Options.HeaderFooterOptions.HeaderLine,
				HeaderFontName: t.Options.HeaderFooterOptions.HeaderFontName,
				HeaderFontSize: t.Options.HeaderFooterOptions.HeaderFontSize,
				HeaderRight:    t.Options.HeaderFooterOptions.HeaderRight,
				HeaderSpacing:  t.Options.HeaderFooterOptions.HeaderSpacing,
			},
		},
	}
}
