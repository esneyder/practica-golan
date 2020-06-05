package bd

import (
	"context"
	"time"

	"github.com/esneyder/practica-golan/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertRegister*/
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("practica-go")
	col := db.Collection("users")
	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
