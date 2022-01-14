package tests

//import (
//	htmlparser "github.com/Revazashvili/easer/htmlparser/usecase"
//	template "github.com/Revazashvili/easer/mongo/usecase"
//	pdf "github.com/Revazashvili/easer/pdf/usecase"
//	"github.com/Revazashvili/easer/template/repository/mongo"
//	"testing"
//)
//
//func TestRenderPdf(t *testing.T) {
//	// add mongo
//	repo := mongo.NewTemplateRepository(options)
//	id, err := repo.AddTemplate(templateSample)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	htmlParser := htmlparser.NewHtmlParser()
//	tempUseCase := template.NewTemplateUseCase(repo, htmlParser)
//	pdfCreator := pdf.NewCreator(htmlParser)
//	pdfRenderer := pdf.NewPdfRenderer(tempUseCase, pdfCreator)
//	// create pdf with added mongo
//	bytes, err := pdfRenderer.Render(id, data)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if len(bytes) < 0 {
//		t.Fatal("rendered pdf byte length is less than zero")
//	}
//	t.Log("Length of bytes %i", len(bytes))
//}
