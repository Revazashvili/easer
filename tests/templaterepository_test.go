package tests

import (
	"github.com/Revazashvili/easer/template/repository/mongo"
	"testing"
)

func TestGetTemplates(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	_, err := tempRepo.AddTemplate(templateSample)
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
	id, err := tempRepo.AddTemplate(templateSample)
	if err != nil {
		t.Fail()
	}
	t.Logf("Inserted templateSample id: %s", id)

	tm, err := tempRepo.GetTemplate(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tm)
}

func TestAddTemplate(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	_, err := tempRepo.AddTemplate(templateSample)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateTemplate(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	id, err := tempRepo.AddTemplate(templateSample)
	if err != nil {
		t.Fail()
	}

	templateSample.Name = "test updated"
	_, err = tempRepo.UpdateTemplate(id, templateSample)
	if err != nil {
		return
	}
}

func TestDeleteTemplate(t *testing.T) {
	tempRepo := mongo.NewTemplateRepository(options)
	id, err := tempRepo.AddTemplate(templateSample)
	if err != nil {
		t.Fail()
	}

	err = tempRepo.DeleteTemplate(id)
	if err != nil {
		t.Fail()
	}
}
