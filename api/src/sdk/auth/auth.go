package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/TaylorCoons/daq-stack/src/models"
	"github.com/TaylorCoons/daq-stack/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// TODO: move these into config
var masterUser string = "admin"
var masterPass string = "pass"
var tokenSize int = 32
var expirySeconds int32 = 60 * 60

var database = "operations"
var collection = "adminTokens"

func IndexTables(client *mongo.Client) {
	collection := client.Database(database).Collection(collection)

	index := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "createdAt", Value: bsonx.Int32(1)}},
		Options: options.Index().SetExpireAfterSeconds(expirySeconds).SetName("expiration"),
	}
	_, err := collection.Indexes().CreateOne(utils.TimeoutCtx(10), index)
	if err != nil {
		panic(err)
	}
}

func BasicAuth(username string, password string) bool {
	return (masterUser == username && masterPass == password)
}

func CreateToken(client *mongo.Client) (models.Token, error) {
	// TODO: Hash token and store into DB
	collection := client.Database(database).Collection(collection)
	token, err := GenerateToken(tokenSize)
	if err != nil {
		return models.Token{}, err
	}
	collection.InsertOne(utils.TimeoutCtx(10), token)

	return token, nil
}

func RenewToken(client *mongo.Client, key string) (models.Token, error) {
	collection := client.Database(database).Collection(collection)
	collection.DeleteOne(utils.TimeoutCtx(10), bson.M{"key": key})
	newToken, err := CreateToken(client)
	if err != nil {
		return models.Token{}, err
	}
	return newToken, nil
}

func ValidateToken(client *mongo.Client, key string) bool {
	collection := client.Database(database).Collection(collection)
	err := collection.FindOne(utils.TimeoutCtx(10), bson.M{"key": key}).Decode(&models.Token{})
	return err == nil
}

func RevokeToken(client *mongo.Client, key string) error {
	collection := client.Database(database).Collection(collection)
	collection.DeleteOne(utils.TimeoutCtx(10), bson.M{"key": key})
	return nil
}

func RevokeAll(client *mongo.Client) error {
	collection := client.Database(database).Collection(collection)
	collection.DeleteMany(utils.TimeoutCtx(10), bson.M{})
	return nil
}

func SupportedAuthTypes() models.SupportedAuth {
	return models.SupportedAuth{
		SupportedAuth: []string{"Basic Authorization"},
	}
}

func generateKey(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func wrapToken(key string) models.Token {
	return models.Token{
		Key:       key,
		CreatedAt: time.Now().UTC(),
		ExpiresOn: time.Now().Add(time.Second * time.Duration(expirySeconds)).UTC(),
	}
}

func GenerateToken(length int) (models.Token, error) {
	key, err := generateKey(length)
	if err != nil {
		return models.Token{}, err
	}
	return wrapToken(key), nil
}
