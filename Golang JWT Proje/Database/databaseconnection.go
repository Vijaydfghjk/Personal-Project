package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {

		log.Fatal("error loading.env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	//mongo.NewClient(options.Client().ApplyURI(MongoDb))

	clientOptions := options.Client().ApplyURI(MongoDb)

	myclient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = myclient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	//myclient.Connect(ctx)

	fmt.Println("connected")

	return myclient
}

var Client *mongo.Client = DBinstance()

func Opencollection(client *mongo.Client, collectionname string) *mongo.Collection{
 

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionname)
	return collection 
}
