package template

import (
	"github.com/Revazashvili/easer/models"
)

type repository interface {
	GetTemplates() ([]*models.Template, error)
	GetTemplate(id string) (*models.Template, error)
	AddTemplate(t *models.Template) (string, error)
	UpdateTemplate(t *models.Template) (string, error)
	DeleteTemplate(id string) error
}