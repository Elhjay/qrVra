package database

import (
	"context"
	"os"

	"main/errorHandlers/standardError"

	"github.com/joho/godotenv"

	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	err := godotenv.Load()
	standardError.StddErr(err, "Error reading env files")

	mongo_url := os.Getenv("MONGO_URL")
	dbOptions := options.Client().ApplyURI(mongo_url)
	client, err = mongo.Connect(context.Background(), dbOptions)
	standardError.StddErr(err, "Error connecting to database")

	fmt.Println("Connected to database")
}

func GetClient() *mongo.Client {
	return client
}
