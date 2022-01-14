package tests

//
//import (
//	"github.com/Revazashvili/easer/models"
//	"github.com/Revazashvili/easer/template/repository/mongo"
//)
//
//var html = `
//<!DOCTYPE html>
//<html lang="en">
//    <head>
//        <meta charset="UTF-8" />
//    </head>
//    <body>
//		<h1>{{.PageTitle}}</h1>
//		<ul>
//			{{range .Todos}}
//				{{if .Done}}
//					<li class="done">{{.Title}}</li>
//				{{else}}
//					<li>{{.Title}}</li>
//				{{end}}
//			{{end}}
//		</ul>
//    </body>
//</html>`
//
//type Todo struct {
//	Title string `json:"name"`
//	Done  bool   `json:"done"`
//}
//
//type TodoPage struct {
//	PageTitle string
//	Todos     []Todo
//}
//
//var data = TodoPage{
//	PageTitle: "My TODO list",
//	Todos: []Todo{
//		{Title: "Task 1", Done: false},
//		{Title: "Task 2", Done: true},
//		{Title: "Task 3", Done: true},
//	},
//}
//
//var options = mongo.DbOptions{
//	Uri:              "mongodb://localhost:27017",
//	DbName:           "template_test",
//	TemplateCollName: "templates",
//}
//
//var templateSample = models.Template{
//	Owner:        "App",
//	TemplateBody: html,
//	Name:         "test",
//	Description:  "test",
//	Options: models.Options{
//		Orientation:          "Portrait",
//		DisableInternalLinks: false,
//		DisableExternalLinks: false,
//		NoBackground:         true,
//		Margin: models.Margin{
//			Top:    1,
//			Right:  1,
//			Left:   1,
//			Bottom: 1,
//		},
//		PrintBackground:     false,
//		NoImages:            false,
//		Grayscale:           false,
//		Format:              "A4",
//		Dpi:                 2,
//		EnableForms:         false,
//		DisplayHeaderFooter: false,
//		HeaderFooterOptions: models.HeaderAndFooterOptions{
//			FooterCenter:   "ad",
//			HeaderFontName: "asd",
//		},
//	},
//}
