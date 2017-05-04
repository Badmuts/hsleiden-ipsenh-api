package api

import (
	"time"
)

type Datapoint struct {
	Sensor    *Sensor   `json:"sensor"`
	Key       int       `json:"key"`
	Value     float32   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type Datapoints struct {
	datapoints map[int]Datapoint
}
