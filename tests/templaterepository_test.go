package tests

import (
	"fmt"
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/template/repository/mongo"
	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func initDB() *mgo.Database {
	uri := viper.GetString("mongo.uri")
	fmt.Println(uri)
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}
	db := session.DB(viper.GetString("mongo.name"))
	return db
}

func TestShouldAddTemplateWithoutError(t *testing.T) {
	db := initDB()
	tempRepo := mongo.NewTemplateRepository(db, "templates")
	template := &models.Template{
		Owner:        "App",
		TemplateBody: "some html",
		Name:         "test",
		Description:  "test",
		Options: &models.Options{
			Orientation:          "A4",
			DisableInternalLinks: false,
			DisableExternalLinks: false,
			NoBackground:         true,
		},
	}
	_, err := tempRepo.AddTemplate(template)
	if err != nil {
		t.Fail()
	}
}

func TestGetTemplatesReturnsNothing(t *testing.T) {
	db := initDB()
	tempRepo := mongo.NewTemplateRepository(db, "templates")
	_, err := tempRepo.GetTemplates()
	if err != nil {
		t.Fatal(err)
		return
	}
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestSample(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}
