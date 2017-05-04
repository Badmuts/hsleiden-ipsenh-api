package api

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// var room Room

	var room = Room{Name: "G3.076", Occupation: 20, MAX_CAPACITY: 15, Size: 20}

	// room.Name = "G3.076"
	// room.Occupation = 20
	// room.MAX_CAPACITY = 15
	// room.Size = 20

	if room.Occupation > room.MAX_CAPACITY {
		fmt.Println(room.Name, room.Occupation, room.MAX_CAPACITY)
		fmt.Println("Too many students in this room")
		t.Failed()
	}
}
