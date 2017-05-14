package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Datapoint struct {
	Sensor    *Sensor         `json:"sensor"`
	SensorIDS []bson.ObjectId `json:"-" bson:"sensors"`
	Key       int             `json:"key"`
	Value     float32         `json:"value"`
	Timestamp int64           `json:"timestamp"`
	db        *mgo.Database
}
