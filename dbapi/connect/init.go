package connectdb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb connection string goes here" // env variables in GO
var dbName = "movies"
var colName = "watchlist"

var collection *mongo.Collection

func ConnectDB() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
} 
