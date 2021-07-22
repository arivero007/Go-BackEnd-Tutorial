package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweetBD(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}
