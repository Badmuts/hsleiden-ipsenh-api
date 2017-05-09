package model

import "gopkg.in/mgo.v2/bson"
import "gopkg.in/mgo.v2"

type Hub struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name    string        `json:"name"`
	sensors []Sensor      `json:"sensors"`
}

type HubJSON struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name    string        `json:"name"`
	Sensors []Sensor      `json:"sensors"`
}

func (h *Hub) Sensors(db *mgo.Database) HubJSON {
	hub := HubJSON{}
	hub.Name = h.Name
	hub.Id = h.Id
	hub.Sensors = GetSensorsByHubId(h.Id.Hex(), db)
	return hub
}
