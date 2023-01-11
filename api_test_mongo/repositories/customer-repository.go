package repositories

import (
	"context"
	"fmt"

	"api_test_crud_mongo/clients"
	"api_test_crud_mongo/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbColl *mongo.Collection

func GetCustomer(id string) (entities.Customer, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Customer{}, err
	}

	result := dbColl.FindOne(context.Background(), bson.M{"_id": objectId})
	var customer entities.Customer
	result.Decode(&customer)
	return customer, nil
}

func InsertCustomer(customer entities.Customer) (entities.Customer, error) {
	customer.ID = primitive.NewObjectID()
	result, err := dbColl.InsertOne(context.Background(), customer)

	if err != nil {
		return entities.Customer{}, err
	}
	customer.ID = result.InsertedID.(primitive.ObjectID)
	return customer, nil
}

func init() {
	dbColl = clients.Database.Collection("customers")
	fmt.Println("Started collection customers")
}
