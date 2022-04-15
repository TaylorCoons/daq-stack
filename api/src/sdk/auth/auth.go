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

func RenewToken(client *mongo.Client, token models.Token) (models.Token, error) {
	if ValidateToken(client, token) {
		return models.Token{}, &NotAuthorized{}
	}
	// Delete token
	collection := client.Database(database).Collection(collection)
	collection.DeleteMany(helpers.TimeoutCtx(10), bson.M{"key": token.Key})
	newToken, err := GenerateToken(tokenSize)
	if err != nil {
		return models.Token{}, err
	}
	collection.InsertOne(helpers.TimeoutCtx(10), newToken)
	return newToken, nil
}

func ValidateToken(client *mongo.Client, token models.Token) bool {
	collection := client.Database(database).Collection(collection)
	res := collection.FindOne(helpers.TimeoutCtx(10), bson.M{"key": token.Key}).Decode(&models.Token{})
	return res != nil
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
