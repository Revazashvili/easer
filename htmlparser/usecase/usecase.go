package htmlparser

import (
	"bytes"
	"github.com/Revazashvili/easer/htmlparser"
	"html/template"
	"log"
)

type HtmlParserUseCase struct {
}

func NewHtmlParser() *HtmlParserUseCase {
	return &HtmlParserUseCase{}
}

func (tp *HtmlParserUseCase) Parse(name string, html string, data interface{}) (string, error) {
	tmpl, err := template.New(name).Parse(html)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return "", htmlparser.ErrParseTemplate
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return "", htmlparser.ErrParseDataToTemplate
	}
	return buf.String(), nil
}
