package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mongoConnector/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// load config
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(config.URI)

	// Create a new client
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the client is connected
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Select a database and collection
	collection := client.Database("testdb").Collection("users")

	err = testit(collection)
	if err != nil {
		log.Fatal(err)
	}

	// // Close the connection when you're done
	// // Defer the disconnection by 5 minutes
	// go func() {
	// 	time.Sleep(time.Duration(config.MongoOpenTime) * time.Minute)
	// 	err := client.Disconnect(context.Background())
	// 	if err != nil {
	// 		log.Printf("Error disconnecting: %v", err)
	// 	} else {
	// 		fmt.Println("Disconnected from MongoDB after 5 minutes")
	// 	}
	// }()

	// // Keep the main function running for demonstration
	// time.Sleep(time.Duration(config.MongoOpenTime+1) * time.Minute)

	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func testit(collection *mongo.Collection) error {

	// Create a document to insert
	newUser := bson.D{
		{"name", "Alice"},
		{"age", 25},
		{"email", "alice@example.com"},
	}

	// Insert the document
	_, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a new user into the database.")

	// Query the document we just inserted
	var result bson.D
	err = collection.FindOne(context.Background(), bson.D{{"name", "Alice"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Found a user: %v\n", result)
	return nil
}
func loadConfig(filename string) (config models.Config, err error) {
	config = models.Config{}

	// Read the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// Unmarshal the JSON into a Config struct

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
