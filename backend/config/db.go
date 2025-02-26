package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, *mongo.Collection) {
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ Erreur de connexion à MongoDB :", err)
	}
	collection := client.Database("cs2-case-tracker").Collection("prices")
	return client, collection
}
