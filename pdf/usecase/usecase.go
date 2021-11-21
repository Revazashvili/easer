package pdf

import (
	"github.com/Revazashvili/easer/pdf"
	"github.com/Revazashvili/easer/template"
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

func (p *UseCase) Render(id string, data interface{}) ([]byte, bool) {
	t, err := p.tu.Find(id)
	if err != nil {
		return nil, true
	}
	return p.pdfCreator.Create(t, data)
}
