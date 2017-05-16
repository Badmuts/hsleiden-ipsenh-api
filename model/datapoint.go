package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Datapoint struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	SensorID  bson.ObjectId `json:"-" bson:"sensor"`
	Key       string        `json:"key" bson:"key"`
	Value     float64       `json:"value" bson:"value"`
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

	d.SensorID = bson.ObjectIdHex("5915a9e7932c2b024d18561e")

	//get sensor and save datapoint to sensor in DB

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

func BulkSaveDatapoints(db *mgo.Database, datapoints []Datapoint) (savedDatapoints []Datapoint, err error) {
	for index, datapoint := range datapoints {
		if datapoint.ID == "" {
			datapoints[index].ID = bson.NewObjectId()
		}
		err = db.C("datapoint").Insert(&datapoints[index])
		if err != nil {
			break
		}
	}

	return datapoints, err

}
