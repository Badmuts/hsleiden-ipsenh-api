package api

type Room struct {
	Name         string `json:"name"`
	Size         int    `json:"size"`
	MAX_CAPACITY int    `json:"max_capacity"`
	Occupation   int    `json:"occupation"`
	Hubs         Hubs   `json:"hubs"`
}

type Rooms struct {
	rooms map[int]Room
}
