package ags

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
 * Get all the items in the specified TQOS database
 */
func GetItems(client *mongo.Client, dbName string) ([]Item, error) {
	var myItems []Item
	collection := client.Database(dbName).Collection("items")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return myItems, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &myItems)
	if err != nil {
		return myItems, err
	}

	return myItems, nil
}

/**
 * Set an external ID that links the internal item to the external product
 */
func UpdateExternalId(client *mongo.Client, dbName string, id string, name string, value string) {
	collection := client.Database(dbName).Collection("items")

	filter := bson.D{{Key: "id", Value: id}}
	updatedField := "externalId." + name
	update := bson.D{{Key: "$set", Value: bson.D{{Key: updatedField, Value: value}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}
