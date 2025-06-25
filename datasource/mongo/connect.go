package mongo

import (
	
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(uri string) (*mongo.Client, error){
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	return client, err
}