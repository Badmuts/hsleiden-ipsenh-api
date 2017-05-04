package main

type Room struct {
	Name         string
	Size         int
	MAX_CAPACITY int
	Occupation   int
	Hubs         []*Hub
}
