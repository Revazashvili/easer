package tests

import (
	htmlparser "github.com/Revazashvili/easer/htmlparser/usecase"
	"testing"
)

func TestHtmlParser(t *testing.T) {
	htmlParser := htmlparser.NewHtmlParser()
	parsedHtml, err := htmlParser.Parse("test", html, data)
	if err != nil {
		t.Fatal(err)
	}

	if len(parsedHtml) < 0 {
		t.Fatal("returned html is empty")
	}
	t.Log(parsedHtml)
}
