package mongo

import (
	"context"
	"errors"
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/template"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ErrGeneral = errors.New("Error")

type DbOptions struct {
	Uri              string
	DbName           string
	TemplateCollName string
}

type TemplateRepository struct {
	dbOptions DbOptions
}

func NewTemplateRepository(options DbOptions) TemplateRepository {
	return TemplateRepository{
		dbOptions: options,
	}
}

func getTemplateCollection(dbOptions DbOptions) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(dbOptions.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, ErrGeneral
	}
	collection := client.Database(dbOptions.DbName).Collection(dbOptions.TemplateCollName)
	return collection, nil
}

func disconnect(coll *mongo.Collection) {
	client := coll.Database().Client()
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

func (tr TemplateRepository) GetTemplates() ([]models.Template, error) {
	coll, err := getTemplateCollection(tr.dbOptions)
	if err != nil {
		return nil, err
	}
	defer disconnect(coll)
	cur, err := coll.Find(
		context.TODO(),
		bson.D{},
	)
	if err != nil {
		log.Fatal(err)
		return nil, template.ErrTemplateNotFound
	}
	var ts []Template
	if err = cur.All(context.TODO(), &ts); err != nil {
		log.Fatal(err)
		return nil, template.ErrTemplateNotDeleted
	}
	return AsDomainList(ts), nil
}

func (tr TemplateRepository) GetTemplate(id string) (models.Template, error) {
	coll, err := getTemplateCollection(tr.dbOptions)
	if err != nil {
		return models.Template{}, err
	}
	defer disconnect(coll)
	var t Template
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		return models.Template{}, template.ErrTemplateNotFound
	}
	err = coll.FindOne(
		context.TODO(),
		bson.M{"_id": bson.M{"$eq": objID}},
	).Decode(&t)
	if err != nil {
		log.Fatal(err)
		return models.Template{}, template.ErrTemplateNotFound
	}
	return AsDomain(t), nil
}

func (tr TemplateRepository) AddTemplate(t models.Template) (string, error) {
	coll, err := getTemplateCollection(tr.dbOptions)
	if err != nil {
		return "", err
	}
	defer disconnect(coll)
	insertResult, err := coll.InsertOne(context.TODO(), AsDbModel(t, primitive.NewObjectID()))
	if err != nil {
		log.Fatal(err)
		return "", template.ErrTemplateNotCreated
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (tr TemplateRepository) UpdateTemplate(id string, t models.Template) (bool, error) {
	coll, err := getTemplateCollection(tr.dbOptions)
	if err != nil {
		return false, err
	}
	defer disconnect(coll)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		return false, template.ErrTemplateNotUpdated
	}
	update := bson.M{
		"$set": AsDbModel(t, objID),
	}
	_, err = coll.UpdateOne(context.TODO(),
		bson.M{"_id": bson.M{"$eq": objID}},
		update)
	if err != nil {
		log.Fatal(err)
		return false, template.ErrTemplateNotUpdated
	}
	return true, nil
}

func (tr TemplateRepository) DeleteTemplate(id string) error {
	coll, err := getTemplateCollection(tr.dbOptions)
	if err != nil {
		return err
	}
	defer disconnect(coll)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		return template.ErrTemplateNotDeleted
	}
	_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}})
	if err != nil {
		log.Fatal(err)
		return template.ErrTemplateNotDeleted
	}
	return nil
}
