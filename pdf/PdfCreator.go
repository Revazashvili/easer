package pdf

import "github.com/Revazashvili/easer/models"

type Creator interface {
	Create(t models.Template, data interface{}) ([]byte, bool)
}
