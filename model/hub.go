package model

type Hub struct {
	Name   string
	Sensor map[int]Sensor `json:"sensors"`
}
