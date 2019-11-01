package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://your_username:your_password@localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("mongo.Connect() ERROR: ", err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("some_database").Collection("some_collection")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	models := []mongo.IndexModel{
		{
			Keys: bson.M{
				"first_name": 1,
			},
		},
		{
			Keys: bson.M{
				"second_name": 1,
			},
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err = col.Indexes().CreateMany(ctx, models, opts)

	if err != nil {
		fmt.Println("Indexes().CreateMany() ERROR:", err)
		os.Exit(1)
	} else {
		fmt.Println("CreateMany() option:", opts)
	}
}
