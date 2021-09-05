package models

type Template struct {
	Id           string
	Name         string
	Description  string
	Owner        string
	TemplateBody string
	Options      *Options
}
