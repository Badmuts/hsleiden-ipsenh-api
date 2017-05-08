package api

type Hub struct {
	Sensor map[int]Sensor `json:"sensors"`
}
