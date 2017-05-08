package api

type Building struct {
	Name     string `json:"name"`
	Rooms    Rooms  `json:"rooms"`
	Location string `json:"location"`
}
