package main

type OccupationSensor struct {
	In_sensor_datapoints  []Datapoint
	Out_sensor_datapoints []Datapoint
	Total_entrances       int
	Total_exits           int
	Current_occupants     int
}
