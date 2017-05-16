package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Datapoint struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	SensorID  bson.ObjectId `json:"sensor_id" bson:"sensor"`
	Key       string        `json:"key" bson:"key"`
	Value     float64       `json:"value" bson:"value"`
	Timestamp int64         `json:"timestamp" bson:"timestamp"`
	DB        *mgo.Database `json:"-" bson:"-"`
}

// HubModel creates a Hub which can be used to query the db
func DatapointModel(db *mgo.Database) *Datapoint {
	datapoint := new(Datapoint)
	datapoint.DB = db
	return datapoint
}

// Save saves d to the database
func (d *Datapoint) Save() (info *mgo.ChangeInfo, err error) {
	if d.ID.Hex() == "" {
		d.ID = bson.NewObjectId()
	}

	info, err = d.DB.C("datapoint").Upsert(d, d)
	if err != nil {
		return info, err
	}

	return info, err
}

// Remove removes d from the database
func (d *Datapoint) Remove(ID string) error {
	return d.DB.C("datapoint").Remove(d)
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
