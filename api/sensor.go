package main

type Sensor struct {
	Name       string
	SensorType string
	Status     bool
	UuId       int
	Datapoints Datapoints
	Hub        *Hub
}

type Sensors struct {
	sensors map[int]Sensor
}
