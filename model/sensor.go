package model

import (
	"log"

	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Sensor struct {
	Id         bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Name       string            `json:"name"`
	SensorType string            `json:"sensorType"`
	Status     bool              `json:"status"`
	UUID       int               `json:"uuid"`
	Datapoints map[int]Datapoint `json:"datapoints"`
	Hub        *Hub              `json:"hub"`
}

func GetSensorsByHubId(id string, db *mgo.Database) []Sensor {
	fmt.Printf("ID " + id)
	sensors := []Sensor{}
	err := db.C("sensor").Find(bson.M{"hub": id}).All(&sensors)

	if err != nil {
		log.Fatal(err)
	}

	return sensors
}
