package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Connect() {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		log.Fatal("❌ MONGODB_URL not set in environment")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("❌ Failed to create Mongo client: %v", err)
	}

	// Ping database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB (Ping failed): %v", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")

	// Assign global variables only after successful ping
	Client = client
	DB = client.Database("social")

	fmt.Println("📦 Database selected: social")
}
