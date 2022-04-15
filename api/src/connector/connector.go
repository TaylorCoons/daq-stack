package connector

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseConnector struct {
	mu     sync.Mutex
	client *mongo.Client
}

var connector DatabaseConnector = DatabaseConnector{}

func Set(client *mongo.Client) {
	connector.mu.Lock()
	defer connector.mu.Unlock()
	connector.client = client
}

func Get() *mongo.Client {
	connector.mu.Lock()
	defer connector.mu.Unlock()
	return connector.client
}
