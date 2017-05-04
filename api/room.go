package api

type Room struct {
	Name         string
	Size         int
	MAX_CAPACITY int
	Occupation   int
	Hubs         Hubs
}

type Rooms struct {
	rooms map[int]Room
}
