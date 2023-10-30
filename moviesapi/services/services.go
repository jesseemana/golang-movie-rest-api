package services

import (
	"fmt"
	"log"
	"context"

	model "gihub.com/jesseemana/movieapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection *mongo.Collection

// Insert a single movie into the database 
func InsertMovie(movie model.Movies) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in database with id:", inserted.insertedID)
}

// Update a movie in the database
func UpdateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

// Delete a single movie from the database
func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count:", deleteCount)
}

// Delete all movies from the database
func DeleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all", deleteResult.DeletedCount, "movies!")

	return deleteResult.DeletedCount
}

// Get all movies from the database
func GetAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}
