package htmlparser

import (
	"bytes"
	"html/template"
	"log"
)

type UseCase struct {
}

func NewHtmlParser() *UseCase {
	return &UseCase{}
}

var emptyString = ""

func (tp *UseCase) Parse(name string, html string, data interface{}) (string, bool) {
	tmpl, err := template.New(name).Parse(html)
	if err != nil {
		log.Printf("%s", err.Error())
		return emptyString, true
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Printf("%s", err.Error())
		return emptyString, true
	}
	return buf.String(), false
}
