package api

type TemperatureSensor struct {
	Temperature_sensor_datapoints Datapoints `json:"temperature_sensor_datapoints"`
	Current_temperature           float32    `json:"current_temperature"`
	Avg_temperature               float32    `json:"avg_temperature"`
}
