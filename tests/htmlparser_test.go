package tests

import (
	"github.com/Revazashvili/easer/parsers"
	"testing"
)

func TestHtmlParser(t *testing.T) {
	htmlParser := parsers.NewHtmlParser()

	parsedHtml, ok := htmlParser.Parse("test", html, data)
	if !ok {
		t.Fatal("can't parse html")
	}
	if len(parsedHtml) < 0 {
		t.Fatal("returned html is empty")
	}
	t.Log(parsedHtml)
}
