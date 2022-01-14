package renderers

import (
	"bytes"
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/parsers"
	htmlToPdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

type PdfRenderer struct {
	parsers.Parser
}

func NewPdfRenderer(r parsers.Parser) Renderer {
	return &PdfRenderer{r}
}

func (pr *PdfRenderer) Render(t models.Template, data interface{}) ([]byte, bool) {
	g := htmlToPdf.NewPDFPreparer()
	setOptions(t.Options, g)
	html, ok := pr.Parser.Parse(t.Id, t.TemplateBody, data)
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
