package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/MagnunAVF/my-links/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_STRING = "mongodb://localhost:27017"
const DB_NAME = "links"
const COLLECTION_NAME = "user"

var db *mongo.Database

// DB connection
func init() {
	fmt.Println("Connecting to db ...")

	fmt.Println("Creating mongo client ...")
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connecting to db ...")
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(DB_NAME)
}

func GetAllUsers() []models.User {
	fmt.Println("DAO - GetAllUsers")

	cur, err := db.Collection(COLLECTION_NAME).Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	var user models.User

	for cur.Next(context.Background()) {
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())

	return users
}

func GetUser(id string) models.User {
	fmt.Println("DAO - GetUser")
	fmt.Println(id)

	var user models.User

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}

	err := db.Collection(COLLECTION_NAME).FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func InsertUser(user models.User) {
	fmt.Println("DAO - InsertUser")
	fmt.Println(user)

	_, err := db.Collection(COLLECTION_NAME).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(id string) {
	fmt.Println("DAO - DeletetUser")
	fmt.Println(id)

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}

	_, err := db.Collection(COLLECTION_NAME).DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(id string, data bson.D) {
	fmt.Println("DAO - UpdateUser")
	fmt.Println(id)

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objID}}

	update := bson.D{
		{"$set", data},
	}
	_, err := db.Collection(COLLECTION_NAME).UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
