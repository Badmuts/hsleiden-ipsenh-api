package api

type Sensor struct {
	Name       string     `json:"name"`
	SensorType string     `json:"sensorType"`
	Status     bool       `json:"status"`
	UuId       int        `json:"uuid"`
	Datapoints Datapoints `json:"datapoints"`
	Hub        *Hub       `json:"hub"`
}

type Sensors struct {
	sensors map[int]Sensor
}
