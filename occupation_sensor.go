package main

import "fmt"

type OccupationSensor struct {
	in_sensor_datapoints  []Datapoint
	out_sensor_datapoints []Datapoint
	total_entrances       int
	total_exits           int
	current_occupants     int
}

func main() {
	fmt.Println("OccupationSensor class")
}
