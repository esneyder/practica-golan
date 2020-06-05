package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User  schema datanase  users*/
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	LastName string             `bson:"lastName" json:"lastName"`
	BirtDate time.Time          `bson:"birtDate" json:"birtDate"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password," json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	Banner   string             `bson:"banner" json:"banner"`
	Bio      string             `bson:"bio" json:"bio"`
}
