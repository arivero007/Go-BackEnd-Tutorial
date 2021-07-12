package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/arivero007/Go-BackEnd-Tutorial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("users")

	var u models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&u)

	u.Password = ""
	if err != nil {
		fmt.Println("Usuario no encontrado" + err.Error())
		return u, err
	}
	return u, nil
}
