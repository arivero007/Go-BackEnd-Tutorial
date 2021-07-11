package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson: "_id, omitempty" json: "id"`
	Name      string             `bson: "_name" json: "name, omitempty"`
	FirstName string             `bson: "_firstName" json: "firstname, omitempty"`
	BornDate  time.Time          `bson: "_bornDate" json: "borndate, omitempty"`
	Email     string             `bson: "_email" json: "email"`
	Password  string             `bson: "_password" json: "password, omitempty"`
	Avatar    string             `bson: "_avatar" json: "avatar, omitempty"`
	Banner    string             `bson: "_banner" json: "banner, omitempty"`
	Biography string             `bson: "_biography" json: "biography, omitempty"`
	Location  string             `bson: "_location" json: "location, omitempty"`
	Website   string             `bson: "_website" json: "website, omitempty"`
}
