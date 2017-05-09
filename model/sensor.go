package model

type Sensor struct {
	Name       string            `json:"name"`
	SensorType string            `json:"sensorType"`
	Status     bool              `json:"status"`
	UUID       int               `json:"uuid"`
	Datapoints map[int]Datapoint `json:"datapoints"`
	Hub        *Hub              `json:"hub"`
}
