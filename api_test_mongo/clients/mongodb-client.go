package clients

import (
	"context"
	"fmt"

	"api_test_crud_mongo/startup"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	*mongo.Database
}

type MongoDBConnect struct {
	*mongo.Client
}

var Database *MongoDatabase
var DBConnect *MongoDBConnect

func connect(connectionString string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))

	if err != nil {
		panic(err)
	}

	DBConnect = &MongoDBConnect{client}
	Database = &MongoDatabase{client.Database("test_api")}

	fmt.Println("Successfully connected")
}

func init() {
	connectionString := startup.AppConfig.ConnectionString
	connect(connectionString)
}
