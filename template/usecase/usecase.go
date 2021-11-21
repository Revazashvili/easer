package template

import (
	"github.com/Revazashvili/easer/htmlparser"
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/template"
)

type UseCase struct {
	templateRepo template.Repository
	htmlParser   htmlparser.UseCase
}

func NewTemplateUseCase(templateRepo template.Repository, htmlparser htmlparser.UseCase) template.UseCase {
	return &UseCase{
		templateRepo: templateRepo,
		htmlParser:   htmlparser,
	}
}

var emptyString = ""

func (t *UseCase) All() ([]models.Template, error) {
	return t.templateRepo.GetTemplates()
}

func (t *UseCase) Find(id string) (models.Template, error) {
	return t.templateRepo.GetTemplate(id)
}

func (t *UseCase) Insert(tm models.Template) (string, error) {
	return t.templateRepo.AddTemplate(tm)
}

func (t *UseCase) Update(id string, tm models.Template) (bool, error) {
	return t.templateRepo.UpdateTemplate(id, tm)
}

func (t *UseCase) Delete(id string) error {
	return t.templateRepo.DeleteTemplate(id)
}

func (t *UseCase) Render(id string, data interface{}) (string, error) {
	temp, err := t.templateRepo.GetTemplate(id)
	if err != nil {
		return emptyString, template.ErrTemplateRender
	}
	parsedHtml, err := t.htmlParser.Parse(temp.Id, temp.TemplateBody, data)
	if err != nil {
		return emptyString, template.ErrTemplateRender
	}
	return parsedHtml, nil
}
