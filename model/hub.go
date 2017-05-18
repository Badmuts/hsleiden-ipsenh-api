package model

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hub struct {
	ID        bson.ObjectId   `json:"id" bson:"_id"`
	Name      string          `json:"name" bson:"name"`
	Serial    string          `json:"serialNumber" bson:"serial"`
	SensorIDS []bson.ObjectId `json:"-" bson:"sensors"`
	Sensors   []Sensor        `json:"sensors,omitempty" bson:"-"` // do not store directly in db
	db        *mgo.Database
}

// HubModel creates a Hub which can be used to query the db
func HubModel(db *mgo.Database) *Hub {
	hub := new(Hub)
	hub.db = db
	return hub
}

// Save saves h to the database
func (h *Hub) Save() (info *mgo.ChangeInfo, err error) {
	if h.ID.Hex() == "" {
		h.ID = bson.NewObjectId()
	}

	// Set HubID for each Sensor
	// @todo this should not be here
	for i, _ := range h.Sensors {
		h.Sensors[i].HubID = h.ID
		h.Sensors[i].ID = bson.NewObjectId()
		h.SensorIDS = append(h.SensorIDS, h.Sensors[i].ID)
		h.db.C("sensor").Upsert(h.Sensors[i], h.Sensors[i])
	}

	info, err = h.db.C("hub").Upsert(h, h)
	if err != nil {
		return info, err
	}

	return info, err
}

// Remove removes h from the database
func (h *Hub) Remove(ID string) error {
	return h.db.C("hub").Remove(h)
}

// Sensors retrieves the associated sensors from this h
func (h *Hub) GetSensors() (sensors []Sensor, err error) {
	sensors = []Sensor{}
	err = h.db.C("sensor").Find(bson.M{"_id": bson.M{"$in": h.SensorIDS}}).All(&sensors)
	h.Sensors = sensors
	return sensors, err
}

// Find finds a list of hubs
func (h *Hub) Find() (hubs []Hub, err error) {
	err = h.db.C("hub").Find(bson.M{}).All(&hubs)

	for index, _ := range hubs {
		hubs[index].db = h.db
	}

	return hubs, err
}

func (h *Hub) FindByID(ID string) (hub Hub, err error) {
	if !bson.IsObjectIdHex(ID) {
		return hub, errors.New("Not a valid id")
	}
	err = h.db.C("hub").FindId(bson.ObjectIdHex(ID)).One(&hub)
	hub.db = h.db
	return hub, err
}
