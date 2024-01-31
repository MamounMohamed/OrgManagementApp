package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var dbName string

func InitMongoDB(connectionString string, dbname string) error {
	clientOptions := options.Client().ApplyURI(connectionString)
	// Set a timeout for the connection to MongoDB
	timeout := 10 * time.Second
	clientOptions.ConnectTimeout = &timeout

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")

	err = CreateDatabaseIfNotExists(dbname)
	if err != nil {
		return err
	}
	return nil
}
func CreateDatabaseIfNotExists(dbName string) error {
	// List all databases to check if the database exists
	databases, err := client.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	// Check if the specified database already exists
	for _, db := range databases {
		if db == dbName {
			fmt.Printf("Database '%s' already exists.\n", dbName)
			return nil
		}
	}

	// If the database does not exist, create it
	err = CreateCollections(dbName, "Users", "Organizations")
	if err != nil {
		return err
	}

	fmt.Printf("Database '%s' created successfully.\n", dbName)
	return nil
}
func CreateCollections(dbName string, collections ...string) error {
	// Create collections if they do not exist
	for _, collection := range collections {
		err := client.Database(dbName).CreateCollection(context.Background(), collection)
		if err != nil {
			return err
		}
		fmt.Printf("Collection '%s' created successfully.\n", collection)
	}

	return nil
}

// GetMongoClient returns the MongoDB client
func GetMongoClient() *mongo.Client {
	return client
}
