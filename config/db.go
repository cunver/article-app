package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoDBConnection() *mongo.Database {
	// Database Config
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	newClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Println("Couldn't created client", err)
	}

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = newClient.Connect(ctx)
	if err != nil {
		log.Println("Couldn't connected with context", err)
	}
	//Cancel context to avoid memory leak
	defer cancel()

	// Ping our db connection
	err = newClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Println("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	// Connect to the database local
	db := newClient.Database("local")
	return db
}
