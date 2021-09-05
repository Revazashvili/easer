package mongo

import (
	"github.com/Revazashvili/easer/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
	"github.com/Revazashvili/easer/template"
)

type TemplateRepository struct {
	db *mgo.Collection
}

func NewTemplateRepository(db *mgo.Database,collection string) *TemplateRepository{
	return &TemplateRepository{
		db: db.C(collection),
	}
}

func(tr TemplateRepository) GetTemplates() ([]*models.Template, error) {
	defer tr.db.Database.Session.Close()
	var ts []*Template
	err := tr.db.Find(bson.M{}).All(&ts)
	if err != nil {
		log.Fatalf("%s",err.Error())
		return nil, template.ErrTemplatesNotFound
	}
	return AsDomainList(ts), nil
}

func(tr TemplateRepository) GetTemplate(id string) (*models.Template, error)  {
	defer tr.db.Database.Session.Close()
	t := new(Template)
	err := tr.db.FindId(bson.ObjectIdHex(id)).One(t)
	if err != nil {
		log.Fatalf("%s",err.Error())
		return nil, template.ErrTemplateNotFound
	}
	return AsDomain(t),nil
}

func(tr TemplateRepository) AddTemplate(t *models.Template) (string, error){
	defer tr.db.Database.Session.Close()
	err:=tr.db.Insert(AsDbModel(t))
	if err != nil {
		log.Fatalf("%s",err.Error())
		return "",template.ErrTemplateNotCreated
	}
	return t.Id,nil
}

func(tr TemplateRepository) UpdateTemplate(t *models.Template) (string, error){
	defer tr.db.Database.Session.Close()
	err := tr.db.UpdateId(bson.ObjectIdHex(t.Id),AsDbModel(t))
	if err != nil {
		log.Fatalf("%s",err.Error())
		return "", template.ErrTemplateNotUpdated
	}
	return t.Id, nil
}

func(tr TemplateRepository) DeleteTemplate(id string) error{
	defer tr.db.Database.Session.Close()
	return tr.db.RemoveId(bson.ObjectIdHex(id))
}