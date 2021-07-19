package bd

import (
	"context"
	"time"

	"github.com/arivero007/Go-BackEnd-Tutorial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(u models.User, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("users")

	register := make(map[string]interface{})

	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.FirstName) > 0 {
		register["firstName"] = u.FirstName
	}

	register["bornDate"] = u.BornDate

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		register["location"] = u.Location
	}

	if len(u.Website) > 0 {
		register["website"] = u.Website
	}

	updateString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{"$eq": objID},
	}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}
	return true, nil
}
