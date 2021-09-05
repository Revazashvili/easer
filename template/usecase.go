package template

import "github.com/Revazashvili/easer/models"

type UseCase interface {
	All() ([]*models.Template, error)
	Find(id string) (*models.Template, error)
	Insert(t *models.Template) (string, error)
	Update(t *models.Template) (string, error)
	Delete(id string) error
}