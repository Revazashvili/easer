package tests

import (
	htmlparser "github.com/Revazashvili/easer/htmlparser/usecase"
	"testing"
)

var html = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
    </head>
    <body>
		<h1>{{.PageTitle}}</h1>
		<ul>
			{{range .Todos}}
				{{if .Done}}
					<li class="done">{{.Title}}</li>
				{{else}}
					<li>{{.Title}}</li>
				{{end}}
			{{end}}
		</ul>
    </body>
</html>`

type Todo struct {
	Title string `json:"name"`
	Done  bool   `json:"done"`
}

type TodoPage struct {
	PageTitle string
	Todos     []Todo
}

var data = TodoPage{
	PageTitle: "My TODO list",
	Todos: []Todo{
		{Title: "Task 1", Done: false},
		{Title: "Task 2", Done: true},
		{Title: "Task 3", Done: true},
	},
}

func TestHtmlParser(t *testing.T) {
	htmlParser := htmlparser.NewHtmlParser()

	parsedHtml, err := htmlParser.Parse("test", html, data)
	if err != nil {
		t.Fatal(err)
	}

	if len(parsedHtml) < 0 {
		t.Fatal("returned html is empty")
	}
}
