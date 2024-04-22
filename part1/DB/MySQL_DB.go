package DB

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func DBinstance() (*mongo.Client, error) {
	err := godotenv.Load("/Users/askarabylkhaiyrov/Desktop/GoLang/assigment3/part1/DB/.env")
	if err != nil {
		return nil, err
	}

	MongoDb := os.Getenv("MONGODBURL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected!!")

	return client, nil
}

var Client *mongo.Client
var err error

func init() {
	Client, err = DBinstance()
	if err != nil {
		log.Fatal(err)
	}
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Cluster0").Collection(collectionName)
	return collection
}
