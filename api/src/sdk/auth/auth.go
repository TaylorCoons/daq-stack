package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/TaylorCoons/daq-stack/src/helpers"
	"github.com/TaylorCoons/daq-stack/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// TODO: move these into config
var username string = "admin"
var password string = "pass"
var tokenSize int = 32
var expirySeconds int32 = 60 * 60

var database = "operations"
var collection = "adminTokens"

func IndexTables(client *mongo.Client) {
	collection := client.Database(database).Collection(collection)
	time.Now().UTC()
	index := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "createdAt", Value: bsonx.Int32(1)}},
		Options: options.Index().SetExpireAfterSeconds(expirySeconds).SetName("expiration"),
	}
	_, err := collection.Indexes().CreateOne(helpers.TimeoutCtx(10), index)
	if err != nil {
		panic(err)
	}
}

func basicAuth(encode string) bool {
	expected := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return encode == expected
}

func CreateToken(client *mongo.Client, encode string) (models.Token, error) {
	if !basicAuth(encode) {
		return models.Token{}, &NotAuthorized{}
	}
	// TODO: Hash token and store into DB
	collection := client.Database(database).Collection(collection)
	token, err := GenerateToken(tokenSize)
	if err != nil {
		return models.Token{}, err
	}
	collection.InsertOne(helpers.TimeoutCtx(10), token)

	return token, nil
}

func RevokeToken(client *mongo.Client, encode string) error {
	if !basicAuth(encode) {
		return &NotAuthorized{}
	}
	// TODO: Remove hashed token from DB
	collection := client.Database(database).Collection(collection)
	collection.DeleteMany(helpers.TimeoutCtx(10), bson.D{})
	return nil
}

func SupportedAuthTypes() models.SupportedAuth {
	return models.SupportedAuth{
		SupportedAuth: []string{"Basic Authorization"},
	}
}

func GenerateToken(length int) (models.Token, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return models.Token{}, err
	}
	return models.Token{
		Token:     base64.URLEncoding.EncodeToString(b),
		CreatedAt: time.Now().UTC(),
		ExpiresOn: time.Now().Add(time.Second * time.Duration(expirySeconds)).UTC(),
	}, nil
}
