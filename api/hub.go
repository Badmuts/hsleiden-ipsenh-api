package api

type Hub struct {
	Sensor Sensors `json:"sensors"`
}

type Hubs struct {
	hubs map[int]Hub
}
