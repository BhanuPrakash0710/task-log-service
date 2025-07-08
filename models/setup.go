package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DatabaseName = "to-do-list-userdb"
var UsercollectionName = "users"
var TaskCollectionName = "tasks"

var mongoClient *mongo.Client

func Setup() {
	//userName, password := config.GetDBConfig()
	ConnectionString := "mongodb://admin:password@localhost:27017"
	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	mongoClient = client
	fmt.Println("✅ Connected to MongoDB")
}
