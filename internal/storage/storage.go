package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// MongoURLStore implements the Storage interface using MongoDB.
type MongoURLStore struct {
	client   *mongo.Client
	database string
	collection string
}

// NewMongoURLStore creates a new instance of MongoURLStore.
func NewMongoURLStore(connectionString, database, collection string) *MongoURLStore {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	return &MongoURLStore{
		client:   client,
		database: database,
		collection: collection,
	}
}

// Save stores the shortened URL and its original URL.
func (store *MongoURLStore) Save(shortURL, originalURL string) {
	collection := store.client.Database(store.database).Collection(store.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{"shortURL": shortURL, "originalURL": originalURL})
	if err != nil {
		log.Printf("Could not insert: %v", err)
	}
}

// Load retrieves the original URL for a given shortened URL.
func (store *MongoURLStore) Load(shortURL string) (string, bool) {
	collection := store.client.Database(store.database).Collection(store.collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result struct {
		OriginalURL string `bson:"originalURL"`
	}
	err := collection.FindOne(ctx, bson.M{"shortURL": shortURL}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", false
		}
		log.Printf("Could not find document: %v", err)
		return "", false
	}

	return result.OriginalURL, true
}
