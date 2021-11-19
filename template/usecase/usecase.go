package usecase

import (
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/template"
)

type TemplateUseCase struct {
	templateRepo template.Repository
}

func NewTemplateUseCase(templateRepo template.Repository) *TemplateUseCase {
	return &TemplateUseCase{
		templateRepo: templateRepo,
	}
}

func(t *TemplateUseCase) All() ([]*models.Template, error)  {
	return t.templateRepo.GetTemplates()
}

func(t *TemplateUseCase) Find(id string) (*models.Template, error) {
	return t.templateRepo.GetTemplate(id)
}

func(t *TemplateUseCase) Insert(tm * models.Template) (string,error) {
	return t.templateRepo.AddTemplate(tm)
}

func(t *TemplateUseCase) Update(tm *models.Template) (string,error) {
	return t.templateRepo.UpdateTemplate(tm)
}

func(t * TemplateUseCase) Delete(id string) error {
	return t.templateRepo.DeleteTemplate(id)
}