package api

type OccupationSensor struct {
	In_sensor_datapoints  Datapoints `json:"in_sensor_datapoints"`
	Out_sensor_datapoints Datapoints `json:"out_sensor_datapoints"`
	Total_entrances       int        `json:"total_entrances"`
	Total_exits           int        `json:"total_exits"`
	Current_occupants     int        `json:"current_occupants"`
}
