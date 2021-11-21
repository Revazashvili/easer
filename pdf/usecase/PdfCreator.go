package pdf

import (
	"bytes"
	"github.com/Revazashvili/easer/htmlparser"
	"github.com/Revazashvili/easer/models"
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

func (c *Creator) Create(t models.Template, data interface{}) ([]byte, bool) {
	g := htmlToPdf.NewPDFPreparer()
	setOptions(t.Options, g)
	html, ok := c.htmlParser.Parse(t.Id, t.TemplateBody, data)
	if !ok {
		return nil, true
	}
	setPageOptions(html, t, g)

	jsonBytes, err := g.ToJSON()
	if err != nil {
		log.Println(err)
		return nil, true
	}
	pdfGenerator, err := htmlToPdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Println(err)
		return nil, true
	}
	err = pdfGenerator.Create()
	if err != nil {
		log.Println(err)
		return nil, true
	}
	return pdfGenerator.Bytes(), false
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
