package pdf

import (
	"bytes"
	"github.com/Revazashvili/easer/htmlparser"
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/pdf"
	htmlToPdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

type Creator struct {
	htmlParser htmlparser.UseCase
}

func NewCreator(htmlParser htmlparser.UseCase) *Creator {
	return &Creator{htmlParser: htmlParser}
}

func (c *Creator) Create(t models.Template, data interface{}) ([]byte, error) {
	g := htmlToPdf.NewPDFPreparer()

	setOptions(t.Options, g)

	html, err := c.htmlParser.Parse(t.Id, t.TemplateBody, data)
	if err != nil {
		log.Fatal(err)
		return nil, pdf.ErrRenderPdf
	}
	setPageOptions(html, t, g)

	jsonBytes, err := g.ToJSON()
	if err != nil {
		log.Fatal(err)
		return nil, pdf.ErrRenderPdf
	}
	pdfGenerator, err := htmlToPdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
		return nil, pdf.ErrRenderPdf
	}
	err = pdfGenerator.Create()
	if err != nil {
		log.Fatal(err)
		return nil, pdf.ErrRenderPdf
	}
	return pdfGenerator.Bytes(), nil
}

func setOptions(options models.Options, g *htmlToPdf.PDFGenerator) {
	g.Dpi.Set(options.Dpi)
	g.Orientation.Set(options.Orientation)
	g.Grayscale.Set(options.Grayscale)
	g.PageSize.Set(options.Format)
	g.MarginBottom.Set(options.Margin.Bottom)
	g.MarginLeft.Set(options.Margin.Left)
	g.MarginRight.Set(options.Margin.Right)
	g.MarginTop.Set(options.Margin.Top)
}

func setPageOptions(html string, t models.Template, g *htmlToPdf.PDFGenerator) {
	pageOptions := htmlToPdf.NewPageOptions() //TODO: need actual implementation
	pageReader := htmlToPdf.NewPageReader(strings.NewReader(html))
	pageReader.PageOptions = pageOptions
	g.AddPage(pageReader)
}
