package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx context.Context
)

func ConnectDB() *mongo.Client {
	//Get env
	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect mongo
	DB_URL := os.Getenv("MONGODB_URI")

	mongoconn, err := mongo.NewClient(options.Client().ApplyURI(DB_URL))
	err = mongoconn.Connect(ctx)

	if err != nil {
		panic(err)
	}

	if err := mongoconn.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	return mongoconn

}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("little_brown_book_shop").Collection(collectionName)
	return collection
}
