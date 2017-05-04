package main

type Hub struct {
	Sensor Sensors
}

type Hubs struct {
	hubs map[int]Hub
}
