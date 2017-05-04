package main

type TemperatureSensor struct {
	temperature_sensor_datapoints []Datapoint
	current_temperature           float32
	avg_temperature               float32
}
