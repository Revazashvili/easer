package models

type Options struct {
	Grayscale            bool
	Dpi                  uint
	DisplayHeaderFooter  bool
	HeaderFooterOptions  HeaderAndFooterOptions
	PrintBackground      bool
	Orientation          string
	Format               string
	EnableForms          bool
	DisableExternalLinks bool
	DisableInternalLinks bool
	NoBackground         bool
	NoImages             bool
	Margin               Margin
}
