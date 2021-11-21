package tests

import (
	htmlparser "github.com/Revazashvili/easer/htmlparser/usecase"
	pdf "github.com/Revazashvili/easer/pdf/usecase"
	"github.com/Revazashvili/easer/template/repository/mongo"
	template "github.com/Revazashvili/easer/template/usecase"
	"testing"
)

func TestRenderPdf(t *testing.T) {
	// add template
	repo := mongo.NewTemplateRepository(options)
	id, err := repo.AddTemplate(templateSample)
	if err != nil {
		t.Fatal(err)
	}

	tempUseCase := template.NewTemplateUseCase(repo)
	htmlParser := htmlparser.NewHtmlParser()
	pdfRenderer := pdf.NewPdfRenderer(tempUseCase, htmlParser)
	// create pdf with added template
	bytes, err := pdfRenderer.Render(id, data)
	if err != nil {
		t.Fatal(err)
	}

	if len(bytes) < 0 {
		t.Fatal("rendered pdf byte length is less than zero")
	}
	t.Log("Length of bytes %i", len(bytes))
}
