package model

type Room struct {
	Name        string      `json:"name"`
	Size        int         `json:"size"`
	maxCapacity int         `json:"maxCapacity"`
	Occupation  int         `json:"occupation"`
	Hubs        map[int]Hub `json:"hubs"`
}
