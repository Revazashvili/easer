package models

type HeaderAndFooterOptions struct {
	FooterCenter   string  `json:"footer_center" bson:"footer_center"`
	FooterFontName string  `json:"footer_font_name" bson:"footer_font_name"`
	FooterFontSize uint    `json:"footer_font_size" bson:"footer_font_size"`
	FooterHTML     string  `json:"footer_html" bson:"footer_html"`
	FooterLeft     string  `json:"footer_left" bson:"footer_left"`
	FooterLine     bool    `json:"footer_line" bson:"footer_line"`
	FooterRight    string  `json:"footer_right" bson:"footer_right"`
	FooterSpacing  float64 `json:"footer_spacing" bson:"footer_spacing"`
	HeaderCenter   string  `json:"header_center" bson:"header_center"`
	HeaderFontName string  `json:"header_font_name" bson:"header_font_name"`
	HeaderFontSize uint    `json:"header_font_size" bson:"header_font_size"`
	HeaderHTML     string  `json:"header_html" bson:"header_html"`
	HeaderLeft     string  `json:"header_left" bson:"header_left"`
	HeaderLine     bool    `json:"header_line" bson:"header_line"`
	HeaderRight    string  `json:"header_right" bson:"header_right"`
	HeaderSpacing  float64 `json:"header_spacing" bson:"header_spacing"`
}