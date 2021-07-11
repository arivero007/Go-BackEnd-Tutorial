package bd

import (
	"context"
	"time"

	"github.com/arivero007/Go-BackEnd-Tutorial/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("users")

	pass, _ := EncriptPassword(u.Password)
	u.Password = pass

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
