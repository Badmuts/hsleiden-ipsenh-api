package model

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hub struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name    string        `json:"name"`
	sensors []Sensor
}

// HubJSON is the JSON representation of a Hub this object is different because Sensors are retrieved 'lazy'
type HubJSON struct {
	Hub
	Sensors []Sensor `json:"sensors"`
}

// Sensors retrieves the sensors of this hub and stores them in sensors
func (h *Hub) Sensors(db *mgo.Database) []Sensor {
	h.sensors = GetSensorsByHubID(h.ID.Hex(), db)
	return h.sensors
}

// GetHubBySensorId returns a pointer to Hub that has a relation with given sensor id
// @todo: Maybe move to dao package
func GetHubBySensorID(ID string, db *mgo.Database) *Hub {
	hub := &Hub{}
	err := db.C("hub").Find(bson.M{"sensors": ID}).One(&hub)

	// @todo: maybe return error to handle later
	if err != nil {
		log.Fatal(err)
	}

	return hub
}
