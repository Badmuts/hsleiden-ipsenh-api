package model

import (
	"encoding/json"
	"errors"

	"time"

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

// Status checks if hub retrieved data from sensors in the last 15 min and write the status: 'online', 'offline' or 'needs attention'
func (h *Hub) Status() string {
	toTime := time.Now().Unix()
	fromTime := toTime - (15 * 60)
	// Query sensor data for last 15 minsD
	dpm := DatapointModel(h.db)
	datapoints := make(map[bson.ObjectId][]Datapoint)
	for _, sensor := range h.SensorIDS {
		datapoints[sensor], _ = dpm.Find(bson.M{
			"sensor": sensor,
			"timestamp": bson.M{
				"$gte": fromTime,
				"$lt":  toTime,
			},
		})
	}

	foundData := 0
	for _, datapoints := range datapoints {
		if len(datapoints) > 0 {
			foundData++
		}
	}

	if foundData == len(h.SensorIDS) {
		// If resultset contains data for all sensors everthing is fine
		return "online"
	} else if foundData > 0 && foundData < len(h.SensorIDS) {
		// If resultset is missing some senor data, hub/sensor needs attention
		return "needs attention"
	}
	// If resultset is empty hub/sensor is offline
	return "offline"
}

func (h *Hub) MarshalJSON() ([]byte, error) {
	type Alias Hub
	return json.Marshal(&struct {
		*Alias
		Status string `json:"status"`
	}{
		Status: h.Status(),
		Alias:  (*Alias)(h),
	})
}
