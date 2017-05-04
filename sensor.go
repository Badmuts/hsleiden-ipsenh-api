package main

import "fmt"

type Sensor struct {
	name       string
	sensorType string
	status     bool
	uuid       int
	datapoints []Datapoint
	hub        Hub
}

func main() {
	fmt.Println("Sensor class")
}
