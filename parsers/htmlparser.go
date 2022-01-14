package parsers

import (
	"bytes"
	"html/template"
	"log"
)

type HtmlParser struct {
}

func NewHtmlParser() Parser {
	return &HtmlParser{}
}

var emptyString = ""

func (p *HtmlParser) Parse(name string, html string, data interface{}) (string, bool) {
	tmpl, err := template.New(name).Parse(html)
	if err != nil {
		log.Printf("%s", err.Error())
		return emptyString, false
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Printf("%s", err.Error())
		return emptyString, false
	}
	return buf.String(), true
}
