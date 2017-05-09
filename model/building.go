package model

type Building struct {
	Name     string       `json:"name"`
	Rooms    map[int]Room `json:"rooms"`
	Location string       `json:"location"`
}
