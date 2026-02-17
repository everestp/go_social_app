package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	mongoURI := os.Getenv("MONGODB_URI")
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err == nil {
		fmt.Println("connection to DB..")
		DB = Client.Database("social")
	} else {
		fmt.Printf("error connect to db:%s\n", err.Error())
	}
}