package model

import (
	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Datapoint struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Sensor    *Sensor       `json:"sensor"`
	SensorID  bson.ObjectId `json:"-" bson:"sensor"`
	Key       int           `json:"key,omitempty" bson:"key"`
	Value     float32       `json:"value,omitempty" bson:"value"`
	Timestamp int64         `json:"timestamp" bson:"timestamp"`
	db        *mgo.Database
}

// HubModel creates a Hub which can be used to query the db
func DatapointModel(db *mgo.Database) *Datapoint {
	datapoint := new(Datapoint)
	datapoint.db = db
	return datapoint
}

// Save saves d to the database
func (d *Datapoint) Save() (info *mgo.ChangeInfo, err error) {
	if d.ID.Hex() == "" {
		d.ID = bson.NewObjectId()
	}

	// Set HubID for each Sensor
	// @todo this should not be here
	d.Sensor.ID = bson.NewObjectId()
	d.SensorID = d.Sensor.ID
	d.db.C("sensor").Upsert(d.Sensor, d.Sensor)

	info, err = d.db.C("datapoint").Upsert(d, d)
	if err != nil {
		return info, err
	}

	return info, err
}

// Remove removes d from the database
func (d *Datapoint) Remove(ID string) error {
	return d.db.C("datapoint").Remove(d)
}

// Datapoint retrieves the associated sensor from this d
func (d *Datapoint) GetSensor() (sensor *Sensor, err error) {
	sensor = &Sensor{}
	err = d.db.C("sensor").Find(bson.M{"_id": bson.ObjectIdHex("fewf")}).One(&sensor)
	d.Sensor = sensor

	log.Printf("datapoint: %s", d.SensorID)
	log.Printf("found sensor: %s", sensor)
	log.Printf("Sensor: %s", d.Sensor)

	return sensor, err
}

// Find finds a list of datapoints
func (d *Datapoint) Find() (datapoints []Datapoint, err error) {
	err = d.db.C("datapoint").Find(bson.M{}).All(&datapoints)

	for index, _ := range datapoints {
		datapoints[index].db = d.db
	}

	return datapoints, err
}

func (d *Datapoint) FindByID(ID string) (datapoint Datapoint, err error) {
	if !bson.IsObjectIdHex(ID) {
		return datapoint, errors.New("Not a valid id")
	}
	err = d.db.C("datapoint").FindId(bson.ObjectIdHex(ID)).One(&datapoint)
	datapoint.db = d.db
	return datapoint, err
}
