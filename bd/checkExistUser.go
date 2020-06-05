package bd

import (
	"context"
	"time"

	"github.com/esneyder/practica-golan/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckExistUser validate  if exist  user from database*/
func CheckExistUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	bd := MongoC.Database("practica-go")
	col := bd.Collection("users")
	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
