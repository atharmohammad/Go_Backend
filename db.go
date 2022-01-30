package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getMongoDbConn() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func getMongoDbCollection(dbname string, collectionName string) (*mongo.Collection, error) {
	client, err := getMongoDbConn()
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(dbname).Collection(collectionName)
	return collection, err
}
