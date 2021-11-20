package tests

import (
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/template/repository/mongo"
	"testing"
)

var options = mongo.DbOptions{
	Uri:              "mongodb://localhost:27017",
	DbName:           "template_test",
	TemplateCollName: "templates",
}

var template = models.Template{
	Owner:        "App",
	TemplateBody: "some html",
	Name:         "test",
	Description:  "test",
	Options: models.Options{
		Orientation:          "Portrait",
		DisableInternalLinks: false,
		DisableExternalLinks: false,
		NoBackground:         true,
		Margin: models.Margin{
			Top:    1,
			Right:  1,
			Left:   1,
			Bottom: 1,
		},
		PrintBackground:     false,
		NoImages:            false,
		Grayscale:           false,
		Format:              "A4",
		Dpi:                 2,
		EnableForms:         false,
		DisplayHeaderFooter: false,
		HeaderFooterOptions: models.HeaderAndFooterOptions{
			FooterCenter:   "ad",
			HeaderFontName: "asd",
		},
	},
}

func TestGetTemplates(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	_, err := tempRepo.AddTemplate(template)
	if err != nil {
		t.Fail()
	}

	ts, err := tempRepo.GetTemplates()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
}

func TestGetTemplate(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	id, err := tempRepo.AddTemplate(template)
	if err != nil {
		t.Fail()
	}
	t.Logf("Inserted template id: %s", id)

	tm, err := tempRepo.GetTemplate(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tm)
}

func TestAddTemplate(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	_, err := tempRepo.AddTemplate(template)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateTemplate(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	id, err := tempRepo.AddTemplate(template)
	if err != nil {
		t.Fail()
	}

	template.Name = "test updated"
	_, err = tempRepo.UpdateTemplate(id, template)
	if err != nil {
		return
	}
}

func TestDeleteTemplate(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	id, err := tempRepo.AddTemplate(template)
	if err != nil {
		t.Fail()
	}

	err = tempRepo.DeleteTemplate(id)
	if err != nil {
		t.Fail()
	}
}
