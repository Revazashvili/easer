package tests

import (
	"testing"
)

func TestGetTemplates(t *testing.T) {
	_, err := storage.Add(templateSample)
	if err != nil {
		t.Fail()
	}

	ts, err := storage.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
}

func TestGetTemplate(t *testing.T) {
	id, err := storage.Add(templateSample)
	if err != nil {
		t.Fail()
	}
	t.Logf("Inserted templateSample id: %s", id)
	tm, err := storage.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tm)
}

func TestAddTemplate(t *testing.T) {
	_, err := storage.Add(templateSample)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateTemplate(t *testing.T) {
	id, err := storage.Add(templateSample)
	if err != nil {
		t.Fail()
	}

	templateSample.Name = "test updated"
	_, err = storage.Update(id, templateSample)
	if err != nil {
		return
	}
}

func TestDeleteTemplate(t *testing.T) {
	id, err := storage.Add(templateSample)
	if err != nil {
		t.Fail()
	}
	err = storage.Delete(id)
	if err != nil {
		t.Fail()
	}
}
