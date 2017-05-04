package api

type Datapoint struct {
	Sensor    *Sensor
	Key       int
	Value     float32
	Timestamp int //dit moet nog een timestamp worden
}

type Datapoints struct {
	datapoints map[int]Datapoint
}
