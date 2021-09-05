package http

import (
	"github.com/Revazashvili/easer/models"
)

type Template struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Owner        string   `json:"owner"`
	TemplateBody string   `json:"template_body"`
	Options      *Options `json:"options"`
}


type HeaderAndFooterOptions struct {
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

type Margin struct {
	Top    uint `json:"top"`
	Bottom uint `json:"bottom"`
	Left   uint `json:"left"`
	Right  uint `json:"right"`
}

type Options struct {
	Grayscale            bool                    `json:"grayscale"`
	Dpi                  uint                    `json:"dpi"`
	DisplayHeaderFooter  bool                    `json:"display_header_footer"`
	HeaderFooterOptions  *HeaderAndFooterOptions `json:"header_footer_options"`
	PrintBackground      bool                    `json:"print_background"`
	Orientation          string                  `json:"orientation"`
	Format               string                  `json:"format"`
	EnableForms          bool                    `json:"enable_forms"`
	DisableExternalLinks bool                    `json:"disable_external_links"`
	DisableInternalLinks bool                    `json:"disable_internal_links"`
	NoBackground         bool                    `json:"no_background"`
	NoImages             bool                    `json:"no_images"`
	Margin               *Margin                 `json:"margin"`
}

func asResponse(t *models.Template) *Template {
	return &Template{
		Id: t.Id,
		Name: t.Name,
		Description: t.Description,
		Owner: t.Owner,
		TemplateBody: t.TemplateBody,
		Options: &Options{
			DisableExternalLinks: t.Options.DisableExternalLinks,
			DisableInternalLinks: t.Options.DisableInternalLinks,
			Dpi: t.Options.Dpi,
			DisplayHeaderFooter: t.Options.DisplayHeaderFooter,
			EnableForms: t.Options.EnableForms,
			Format: t.Options.Format,
			Grayscale: t.Options.Grayscale,
			NoImages: t.Options.NoImages,
			NoBackground: t.Options.NoBackground,
			Orientation: t.Options.Orientation,
			PrintBackground: t.Options.PrintBackground,
			Margin: &Margin{
				Top: t.Options.Margin.Top,
				Bottom: t.Options.Margin.Bottom,
				Left: t.Options.Margin.Left,
				Right: t.Options.Margin.Left,
			},
			HeaderFooterOptions: &HeaderAndFooterOptions{
				FooterCenter: t.Options.HeaderFooterOptions.FooterCenter,
				FooterHTML: t.Options.HeaderFooterOptions.FooterHTML,
				FooterLeft: t.Options.HeaderFooterOptions.FooterLeft,
				FooterLine: t.Options.HeaderFooterOptions.FooterLine,
				FooterFontName: t.Options.HeaderFooterOptions.FooterFontName,
				FooterFontSize: t.Options.HeaderFooterOptions.FooterFontSize,
				FooterRight: t.Options.HeaderFooterOptions.FooterRight,
				FooterSpacing: t.Options.HeaderFooterOptions.FooterSpacing,
				HeaderCenter: t.Options.HeaderFooterOptions.HeaderCenter,
				HeaderHTML: t.Options.HeaderFooterOptions.HeaderHTML,
				HeaderLeft: t.Options.HeaderFooterOptions.HeaderLeft,
				HeaderLine: t.Options.HeaderFooterOptions.HeaderLine,
				HeaderFontName: t.Options.HeaderFooterOptions.HeaderFontName,
				HeaderFontSize: t.Options.HeaderFooterOptions.HeaderFontSize,
				HeaderRight: t.Options.HeaderFooterOptions.HeaderRight,
				HeaderSpacing: t.Options.HeaderFooterOptions.HeaderSpacing,
			},
		},
	}
}

func asResponses(ts []*models.Template) []*Template {
	out := make([]*Template,len(ts))
	for i,t := range ts{
		out[i] = asResponse(t)
	}
	return out
}

func asDomain(t *Template) *models.Template {
	return &models.Template{
		Id: t.Id,
		Name: t.Name,
		Description: t.Description,
		Owner: t.Owner,
		TemplateBody: t.TemplateBody,
		Options: &models.Options{
			DisableExternalLinks: t.Options.DisableExternalLinks,
			DisableInternalLinks: t.Options.DisableInternalLinks,
			Dpi: t.Options.Dpi,
			DisplayHeaderFooter: t.Options.DisplayHeaderFooter,
			EnableForms: t.Options.EnableForms,
			Format: t.Options.Format,
			Grayscale: t.Options.Grayscale,
			NoImages: t.Options.NoImages,
			NoBackground: t.Options.NoBackground,
			Orientation: t.Options.Orientation,
			PrintBackground: t.Options.PrintBackground,
			Margin: &models.Margin{
				Top: t.Options.Margin.Top,
				Bottom: t.Options.Margin.Bottom,
				Left: t.Options.Margin.Left,
				Right: t.Options.Margin.Left,
			},
			HeaderFooterOptions: &models.HeaderAndFooterOptions{
				FooterCenter: t.Options.HeaderFooterOptions.FooterCenter,
				FooterHTML: t.Options.HeaderFooterOptions.FooterHTML,
				FooterLeft: t.Options.HeaderFooterOptions.FooterLeft,
				FooterLine: t.Options.HeaderFooterOptions.FooterLine,
				FooterFontName: t.Options.HeaderFooterOptions.FooterFontName,
				FooterFontSize: t.Options.HeaderFooterOptions.FooterFontSize,
				FooterRight: t.Options.HeaderFooterOptions.FooterRight,
				FooterSpacing: t.Options.HeaderFooterOptions.FooterSpacing,
				HeaderCenter: t.Options.HeaderFooterOptions.HeaderCenter,
				HeaderHTML: t.Options.HeaderFooterOptions.HeaderHTML,
				HeaderLeft: t.Options.HeaderFooterOptions.HeaderLeft,
				HeaderLine: t.Options.HeaderFooterOptions.HeaderLine,
				HeaderFontName: t.Options.HeaderFooterOptions.HeaderFontName,
				HeaderFontSize: t.Options.HeaderFooterOptions.HeaderFontSize,
				HeaderRight: t.Options.HeaderFooterOptions.HeaderRight,
				HeaderSpacing: t.Options.HeaderFooterOptions.HeaderSpacing,
			},
		},
	}
}
