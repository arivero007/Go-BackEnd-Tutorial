package bd

import (
	"context"

	"github.com/arivero007/Go-BackEnd-Tutorial/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://arivero007:wAx3DSD9HJkG2MOn@cluster0.zunfv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	logs.InitLogger()
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logs.LogError("DB connection error: ", err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logs.LogError("DB ping error: ", err)
		return client
	}

	logs.LogInfo("Conexion Exitosa con la BD")
	return client
}

func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		logs.LogError("DB checking ping error: ", err)
		return 0
	}
	return 1
}
