package mongo

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Template struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Description  string        `bson:"description"`
	Owner        string        `bson:"owner"`
	TemplateBody string        `bson:"template_body"`
	Options      *Options      `bson:"options"`
	Created      time.Time     `bson:"created"`
	CreatedBy    string        `bson:"created_by"`
	Updated      time.Time     `bson:"updated"`
	UpdatedBy    string        `bson:"updated_by"`
}

type Options struct {
	Grayscale            bool                   `json:"grayscale"`
	Dpi                  uint                   `json:"dpi"`
	DisplayHeaderFooter  bool                   `bson:"display_header_footer"`
	HeaderFooterOptions  *HeaderAndFooterOptions `bson:"header_footer_options"`
	PrintBackground      bool                   `bson:"print_background"`
	Orientation          string                 `bson:"orientation"`
	Format               string                 `bson:"format"`
	EnableForms          bool                   `bson:"enable_forms"`
	DisableExternalLinks bool                   `bson:"disable_external_links"`
	DisableInternalLinks bool                   `bson:"disable_internal_links"`
	NoBackground         bool                   `bson:"no_background"`
	NoImages             bool                   `bson:"no_images"`
	Margin               *Margin                 `bson:"margin"`
}

type Margin struct {
	Top    uint `bson:"top"`
	Bottom uint `bson:"bottom"`
	Left   uint `bson:"left"`
	Right  uint `bson:"right"`
}

type HeaderAndFooterOptions struct {
	FooterCenter   string  `bson:"footer_center"`
	FooterFontName string  `bson:"footer_font_name"`
	FooterFontSize uint    `bson:"footer_font_size"`
	FooterHTML     string  `bson:"footer_html"`
	FooterLeft     string  `bson:"footer_left"`
	FooterLine     bool    `bson:"footer_line"`
	FooterRight    string  `bson:"footer_right"`
	FooterSpacing  float64 `bson:"footer_spacing"`
	HeaderCenter   string  `bson:"header_center"`
	HeaderFontName string  `bson:"header_font_name"`
	HeaderFontSize uint    `bson:"header_font_size"`
	HeaderHTML     string  `bson:"header_html"`
	HeaderLeft     string  `bson:"header_left"`
	HeaderLine     bool    `bson:"header_line"`
	HeaderRight    string  `bson:"header_right"`
	HeaderSpacing  float64 `bson:"header_spacing"`
}