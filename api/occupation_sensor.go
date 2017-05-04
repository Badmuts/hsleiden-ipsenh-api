package api

type OccupationSensor struct {
	In_sensor_datapoints  Datapoints
	Out_sensor_datapoints Datapoints
	Total_entrances       int
	Total_exits           int
	Current_occupants     int
}
