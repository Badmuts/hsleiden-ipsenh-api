package model

import "gopkg.in/mgo.v2/bson"

type Hub struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name"`
	Sensor []Sensor      `json:"sensors"`
}
