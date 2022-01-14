package mongo

import (
	"context"
	"errors"
	"github.com/Revazashvili/easer/models"
	"github.com/Revazashvili/easer/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ErrGeneral = errors.New("error occurred")

type DbOptions struct {
	Uri              string
	DbName           string
	TemplateCollName string
}

type TemplateStorage struct {
	dbOptions DbOptions
}

func NewTemplateStorage(options DbOptions) storage.Storage {
	return &TemplateStorage{
		dbOptions: options,
	}
}

func getTemplateCollection(dbOptions DbOptions) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(dbOptions.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return nil, ErrGeneral
	}
	collection := client.Database(dbOptions.DbName).Collection(dbOptions.TemplateCollName)
	return collection, nil
}

func disconnect(coll *mongo.Collection) {
	client := coll.Database().Client()
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}
}

func (storage *TemplateStorage) GetAll() ([]models.Template, error) {
	coll, err := getTemplateCollection(storage.dbOptions)
	if err != nil {
		return nil, err
	}
	defer disconnect(coll)
	cur, err := coll.Find(
		context.TODO(),
		bson.D{},
	)
	if err != nil {
		log.Println(err)
		return nil, ErrGeneral
	}
	var ts []Template
	if err = cur.All(context.TODO(), &ts); err != nil {
		log.Println(err)
		return nil, ErrGeneral
	}
	return AsDomainList(ts), nil
}

func (storage *TemplateStorage) Get(id string) (models.Template, error) {
	coll, err := getTemplateCollection(storage.dbOptions)
	if err != nil {
		return models.Template{}, err
	}
	defer disconnect(coll)
	var tm Template
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return models.Template{}, ErrGeneral
	}
	err = coll.FindOne(
		context.TODO(),
		bson.M{"_id": bson.M{"$eq": objID}},
	).Decode(&tm)
	if err != nil {
		log.Println(err)
		return models.Template{}, ErrGeneral
	}
	return AsDomain(tm), nil
}

func (storage TemplateStorage) Add(t models.Template) (string, error) {
	coll, err := getTemplateCollection(storage.dbOptions)
	if err != nil {
		return "", err
	}
	defer disconnect(coll)
	insertResult, err := coll.InsertOne(context.TODO(), AsDbModel(t, primitive.NewObjectID()))
	if err != nil {
		log.Println(err)
		return "", ErrGeneral
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (storage TemplateStorage) Update(id string, t models.Template) (bool, error) {
	coll, err := getTemplateCollection(storage.dbOptions)
	if err != nil {
		return false, err
	}
	defer disconnect(coll)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return false, ErrGeneral
	}
	update := bson.M{
		"$set": AsDbModel(t, objID),
	}
	_, err = coll.UpdateOne(context.TODO(),
		bson.M{"_id": bson.M{"$eq": objID}},
		update)
	if err != nil {
		log.Println(err)
		return false, ErrGeneral
	}
	return true, nil
}

func (storage TemplateStorage) Delete(id string) error {
	coll, err := getTemplateCollection(storage.dbOptions)
	if err != nil {

		return ErrGeneral
	}
	defer disconnect(coll)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return ErrGeneral
	}
	_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objID}})
	if err != nil {
		log.Println(err)
		return ErrGeneral
	}
	return nil
}
