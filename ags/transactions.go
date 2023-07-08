package ags

import (
	"context"
	"log"

	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const TransactionDatabaseName = "transactions"
const TransactionCollectionName = "transactions"

/**
 * Get the most recent transaction. Often used to find out when we last retrieved data.
 */
func GetMostRecentTransaction(client *mongo.Client) Transaction {
	collection := client.Database(TransactionDatabaseName).Collection(TransactionCollectionName)
	sort := bson.D{{Key: "sort", Value: bson.D{{Key: "created_at", Value: -1}}}}

	var result Transaction
	err := collection.FindOne(context.TODO(), sort).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

/**
 * Set an external ID that links the internal item to the external product
 */
func InsertTransaction(client *mongo.Client, txn *Transaction) bool {
	collection := client.Database(TransactionDatabaseName).Collection(TransactionCollectionName)
	filter := bson.D{{Key: "identifier", Value: bson.D{{Key: "id", Value: txn.Identifier.Id}}}}
	update := bson.D{{Key: "$setOnInsert", Value: txn}}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		// We expect some to fail due to duplicate key errors
		if !mongo.IsDuplicateKeyError(err) {
			panic(err)
		}
		return false
	}
	return true
}
