package renderers

import "github.com/Revazashvili/easer/models"

type Renderer interface {
	Render(t models.Template, data interface{}) ([]byte, bool)
}
