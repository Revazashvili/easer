package pdf

import (
	"github.com/Revazashvili/easer/pdf"
	"github.com/Revazashvili/easer/template"
	"log"
)

type UseCase struct {
	tu         template.UseCase
	pdfCreator pdf.Creator
}

func NewPdfRenderer(tu template.UseCase, pg pdf.Creator) *UseCase {
	return &UseCase{
		tu:         tu,
		pdfCreator: pg,
	}
}

func (p *UseCase) Render(id string, data interface{}) ([]byte, error) {
	t, err := p.tu.Find(id)
	if err != nil {
		log.Fatalf("Error occured while retreiving template to generate pdf %s", err.Error())
		return nil, pdf.ErrRenderPdf
	}
	return p.pdfCreator.Create(t, data)
}
