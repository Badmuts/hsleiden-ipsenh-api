package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hub struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name    string        `json:"name"`
	sensors []Sensor      `json:"sensors"`
}

type HubJSON struct {
	Hub
	Sensors []Sensor `json:"sensors"`
}

func (h *Hub) Sensors(db *mgo.Database) []Sensor {
	if h.sensors == nil {
		return GetSensorsByHubId(h.Id.Hex(), db)
	}
	return h.sensors
}
