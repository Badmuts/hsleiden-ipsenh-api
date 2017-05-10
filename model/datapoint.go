package model

type Datapoint struct {
	Sensor    *Sensor `json:"sensor"`
	Key       int     `json:"key"`
	Value     float32 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}
