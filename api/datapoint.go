package main

type Datapoint struct {
	Sensor    *Sensor
	Key       int
	Value     float32
	Timestamp int //dit moet nog een timestamp worden
}
