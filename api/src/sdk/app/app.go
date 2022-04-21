package app

import (
	"encoding/base64"
	"time"

	"github.com/TaylorCoons/daq-stack/src/models"
	"github.com/TaylorCoons/daq-stack/src/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Replace this with userID
var database string = "user"
var collection string = "apps"

func CreateApp(client *mongo.Client, app models.App) (models.App, error) {
	collection := client.Database(database).Collection(collection)
	id, err := generateId()
	if err != nil {
		return models.App{}, err
	}
	app.Id = id
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()

	_, err = collection.InsertOne(utils.TimeoutCtx(10), app)
	if err != nil {
		return models.App{}, err
	}
	return app, nil
}

func UpdateApp(client *mongo.Client, new models.App, id string) (models.App, error) {
	collection := client.Database(database).Collection(collection)
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	res := collection.FindOneAndUpdate(
		utils.TimeoutCtx(10),
		bson.M{"id": id},
		bson.M{
			"$set": bson.M{
				"description": new.Description,
				"updatedAt":   time.Now(),
			},
		},
		&opts,
	)
	if res.Err() == mongo.ErrNoDocuments {
		return models.App{}, AppNotFoundError{}
	}
	updated := models.App{}
	err := res.Decode(&updated)
	if err != nil {
		return models.App{}, err
	}
	return updated, nil
}

func ListApps(client *mongo.Client) ([]models.App, error) {
	collection := client.Database(database).Collection(collection)
	cur, err := collection.Find(utils.TimeoutCtx(10), bson.M{})
	if err != nil {
		return nil, err
	}
	list := []models.App{}
	err = cur.All(utils.TimeoutCtx(10), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetApp(client *mongo.Client, id string) (models.App, error) {
	collection := client.Database(database).Collection(collection)
	res := collection.FindOne(utils.TimeoutCtx(10), bson.M{"id": id})
	if res.Err() == mongo.ErrNoDocuments {
		return models.App{}, AppNotFoundError{}
	}
	app := models.App{}
	err := res.Decode(&app)
	if err != nil {
		return models.App{}, err
	}
	return app, nil
}

func DeleteApp(client *mongo.Client, id string) error {
	collection := client.Database(database).Collection(collection)
	res, err := collection.DeleteOne(utils.TimeoutCtx(10), bson.M{"id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return AppNotFoundError{}
	}
	return nil
}

func generateId() (string, error) {
	idUuid, err := uuid.New().MarshalText()
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(idUuid), nil
}
