package model

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hub struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name"`
	Serial  string        `json:"serialNumber"`
	sensors []Sensor
	db      *mgo.Database
}

// HubJSON is the JSON representation of a Hub this object is different because Sensors are retrieved 'lazy'
type HubJSON struct {
	*Hub
	Sensors []Sensor `json:"sensors"`
}

// Save saves h to the database
func (h *Hub) Save(db *mgo.Database) (info *mgo.ChangeInfo, err error) {
	if h.ID.Hex() == "" {
		h.ID = bson.NewObjectId()
	}

	return db.C("hub").Upsert(h, h)
}

// Remove removes h from the database
func (h *Hub) Remove() error {
	return h.db.C("hub").Remove(h)
}

// Sensors retrieves the associated sensors from this h
func (h *Hub) Sensors(db *mgo.Database) (sensors []Sensor, err error) {
	sensors = []Sensor{}
	err = db.C("sensor").Find(bson.M{"hub": h.ID.Hex()}).All(&sensors)
	return sensors, err
}

// JSON creates a HubJSON struct where relationships are also exposed
func (h *Hub) JSON(db *mgo.Database) (hJSON HubJSON, err error) {
	h.sensors, err = h.Sensors(db)

	return HubJSON{h, h.sensors}, err
}

// HubModel creates a Hub which can be used to query the db
func HubModel(db *mgo.Database) *Hub {
	hub := &Hub{}
	hub.db = db
	return &Hub{}
}

// Find finds a list of hubs
func (h *Hub) Find(db *mgo.Database) (hubs []Hub, err error) {
	err = db.C("hub").Find(bson.M{}).Limit(25).All(&hubs)
	return hubs, err
}

func (h *Hub) FindByID(db *mgo.Database, ID string) (hub Hub, err error) {
	if !bson.IsObjectIdHex(ID) {
		return hub, errors.New("Not a valid objectId")
	}
	err = db.C("hub").FindId(bson.ObjectIdHex(ID)).One(&hub)
	return hub, err
}
