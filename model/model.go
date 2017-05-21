package model

import "gopkg.in/mgo.v2/bson"
import "gopkg.in/mgo.v2"

type Model struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	db *mgo.Database
}
