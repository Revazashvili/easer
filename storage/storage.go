package storage

import "github.com/Revazashvili/easer/models"

type Storage interface {
	GetAll() ([]models.Template, error)
	Get(id string) (models.Template, error)
	Add(t models.Template) (string, error)
	Update(id string, t models.Template) (bool, error)
	Delete(id string) error
}
