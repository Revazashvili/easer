package models

type Options struct {
	Grayscale            bool                   `json:"grayscale" json:"grayscale"`
	Dpi                  uint                   `json:"dpi" json:"dpi"`
	DisplayHeaderFooter  bool                   `json:"display_header_footer" bson:"display_header_footer"`
	HeaderFooterOptions  HeaderAndFooterOptions `json:"header_footer_options" bson:"header_footer_options"`
	PrintBackground      bool                   `json:"print_background" bson:"print_background"`
	Orientation          string                 `json:"orientation" bson:"orientation"`
	Format               string                 `json:"format" bson:"format"`
	EnableForms          bool                   `json:"enable_forms" bson:"enable_forms"`
	DisableExternalLinks bool                   `json:"disable_external_links" bson:"disable_external_links"`
	DisableInternalLinks bool                   `json:"disable_internal_links" bson:"disable_internal_links"`
	NoBackground         bool                   `json:"no_background" bson:"no_background"`
	NoImages             bool                   `json:"no_images" bson:"no_images"`
	Margin               Margin                 `json:"margin" bson:"margin"`
}
