package model

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hub struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Serial  string        `json:"serialNumber" bson:"serial"`
	sensors []Sensor
	db      *mgo.Database
}

// HubJSON is the JSON representation of a Hub this object is different because Sensors are retrieved 'lazy'
type HubJSON struct {
	*Hub
	Sensors []Sensor `json:"sensors"`
}

// Save saves h to the database
func (h *Hub) Save() (info *mgo.ChangeInfo, err error) {
	if h.ID.Hex() == "" {
		h.ID = bson.NewObjectId()
	}

	return h.db.C("hub").Upsert(h, h)
}

// Remove removes h from the database
func (h *Hub) Remove(ID string) error {
	return h.db.C("hub").Remove(h)
}

// Sensors retrieves the associated sensors from this h
func (h *Hub) Sensors() (sensors []Sensor, err error) {
	err = h.db.C("sensor").Find(bson.M{"hub": h.ID.Hex()}).All(&sensors)
	h.sensors = sensors
	return sensors, err
}

// JSON creates a HubJSON struct where relationships are also exposed
func (h *Hub) JSON() (hJSON HubJSON, err error) {
	h.sensors, err = h.Sensors()

	return HubJSON{h, h.sensors}, err
}

// HubModel creates a Hub which can be used to query the db
func HubModel(db *mgo.Database) *Hub {
	hub := new(Hub)
	hub.db = db
	return hub
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
		return hub, errors.New("Not a valid objectId")
	}
	err = h.db.C("hub").FindId(bson.ObjectIdHex(ID)).One(&hub)
	return hub, err
}
