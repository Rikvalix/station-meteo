package config

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"time"
)

// ConnectDb Connection à la base de données MongoDB
func ConnectDb() *mongo.Client {
	logger := NewLogger()
	client, err := mongo.Connect(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		logger.Fatalf("Erreur de connexion à MongoDB : %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal("❌ Erreur de connection à MongoDB:", err)
	}

	logger.Info("Connected to MongoDB")
	return client
}

// DB Instance du client
var DB *mongo.Client = ConnectDb()

// GetCollection Renvoi les collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
