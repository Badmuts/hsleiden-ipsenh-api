package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Sensor struct {
	ID         bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name,omitempty" bson:"name"`
	SensorType string        `json:"sensorType,omitempty" bson:"sensorType"`
	Status     bool          `json:"status,omitempty" bson:"status"`
	UUID       int           `json:"UUID,omitempty" bson:"UUID"`
	Gpio_trigger       int   `json:"gpio_trigger,omitempty" bson:"gpio_trigger"`
	Gpio_echo       int   		`json:"gpio_echo,omitempty" bson:"gpio_echo"`
	Datapoints []Datapoint   `json:"datapoints,omitempty" bson:"datapoints"`
	HubID      bson.ObjectId `json:"-" bson:"hub"`
	hub        *Hub
}

// SensorJSON is a wrapper to expose the relation with hub
type SensorJSON struct {
	Sensor
	Hub *Hub `json:"hub"`
}
