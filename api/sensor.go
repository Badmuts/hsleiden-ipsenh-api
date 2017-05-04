package main

type Sensor struct {
	Name       string
	SensorType string
	Status     bool
	UuId       int
	Datapoints []Datapoint
	Hub        *Hub
}
