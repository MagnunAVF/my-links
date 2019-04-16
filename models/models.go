package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Links []Link             `json:"links,omitempty" bson:"links,omitempty"`
}

type Link struct {
	Title string `json:"title,omitempty bson:"title,omitempty""`
	Url   string `json:"url,omitempty bson:"_url,omitempty""`
}
