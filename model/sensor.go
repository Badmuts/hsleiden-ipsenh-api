package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Sensor struct {
	ID         bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Name       string            `json:"name"`
	SensorType string            `json:"sensorType"`
	Status     bool              `json:"status"`
	UUID       int               `json:"uuid"`
	Datapoints map[int]Datapoint `json:"datapoints"`
	hub        *Hub
}

// SensorJSON is a wrapper to expose the relation with hub
type SensorJSON struct {
	Sensor
	Hub *Hub `json:"hub"`
}

// Hub returns the Hub that has a relationship with this sensor.
// func (s *Sensor) Hub(db *mgo.Database) *Hub {
// 	s.hub = GetHubBySensorID(s.ID.Hex(), db)
// 	return s.hub
// }
