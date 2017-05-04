package main

type Sensor struct {
	name       string
	sensorType string
	status     bool
	uuid       int
	datapoints []Datapoint
	hub        Hub
}
