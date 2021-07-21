package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/arivero007/Go-BackEnd-Tutorial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page int64) ([]*models.GetTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db := MongoCN.Database("gotutorial")
	col := db.Collection("tweet")

	var result []*models.GetTweets

	condition := bson.M{
		"userid": ID,
	}

	opt := options.Find()
	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	opt.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opt)

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var register models.GetTweets
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}
