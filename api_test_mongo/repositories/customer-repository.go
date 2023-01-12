package repositories

import (
	"context"
	"errors"
	"fmt"

	"api_test_crud_mongo/clients"
	"api_test_crud_mongo/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func GetCustomer(id string) (entities.Customer, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Customer{}, err
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.M{"_id": objectId},
				bson.M{"active": true},
			}},
	}

	result := collection.FindOne(context.Background(), filter)
	var customer entities.Customer
	result.Decode(&customer)
	return customer, nil
}

func InsertCustomer(customer entities.Customer) (entities.Customer, error) {
	customer.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), customer)

	if err != nil {
		return entities.Customer{}, err
	}

	return customer, nil
}

func RaplaceCustomer(customer entities.Customer) (entities.Customer, error) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.M{"_id": customer.ID},
				bson.M{"active": true},
			}},
	}

	result, err := collection.ReplaceOne(context.TODO(), filter, customer)
	if err != nil {
		return entities.Customer{}, err
	}

	if result.MatchedCount == 0 {
		return entities.Customer{}, nil
	}

	return customer, nil
}

func DeleteCustomer(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{
		{"$and",
			bson.A{
				bson.M{"_id": objectId},
				bson.M{"active": true},
			}},
	}

	update := bson.D{{"$set", bson.M{"active": false}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("Document not exist")
	}

	return nil
}

func init() {
	collection = clients.Database.Collection("customers")
	fmt.Println("Started collection customers")
}
