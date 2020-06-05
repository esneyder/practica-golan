package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC  objeto conexion con la base de datos*/
var MongoC = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/practica-go")

/*ConnectBD funcion que me  permite conectar con la base de datos */
func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con  la  BD")
	return client
}

/*CheckConnect chequea la conexión con la base de datos */
func CheckConnect() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
