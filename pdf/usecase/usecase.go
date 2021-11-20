package usecase

import (
	"github.com/Revazashvili/easer/htmlparser"
	"github.com/Revazashvili/easer/template"
	"log"
)

type PdfUseCase struct {
	tu template.UseCase
	hp htmlparser.UseCase
}

func NewPdfUseCase(tu template.UseCase,hp htmlparser.UseCase) *PdfUseCase {
	return &PdfUseCase{
		tu: tu,
		hp: hp,
	}
}

func(p *PdfUseCase) Render(id string,data interface{}) ([]byte,error) {
	t,err := p.tu.Find(id)
	if err!=nil{
		log.Fatalf("Error occured while retreiving template to generate pdf %s",err.Error())
		return nil, err
	}
	s,err := p.hp.Parse(t.Id,t.TemplateBody,data)
	if err!=nil {
		log.Fatalf("Error occured while generating pdf %s", err.Error())
		return nil, err
	}
	return []byte(s), nil
}
