package template

import (
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/template"
)

type UseCase struct {
	templateRepo template.Repository
}

func NewTemplateUseCase(templateRepo template.Repository) template.UseCase {
	return &UseCase{
		templateRepo: templateRepo,
	}
}

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
