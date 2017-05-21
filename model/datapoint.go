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
	DB        *mgo.Database
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
