package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User represents the structure of a user in Go
type User struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Email  string        `json:"email" bson:"email"`
	Age    int           `json:"age" bson:"age"`
	Gender string        `json:"gender" bson:"gender"`
}
