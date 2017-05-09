package api

import (
	"fmt"
	"testing"
)

type Room struct {
	Name         string `json:"name"`
	Size         int    `json:"size"`
	MAX_CAPACITY int    `json:"max_capacity"`
	Occupation   int    `json:"occupation"`
}

func TestCheckIfRoomIsOverOccpupied(t *testing.T) {
	room := Room{Name: "G3.076", Occupation: 20, MAX_CAPACITY: 30, Size: 20}

	if room.Occupation > room.MAX_CAPACITY {
		fmt.Println("Name: ", room.Name, "Occupation: ", room.Occupation, "Max Capacity: ", room.MAX_CAPACITY)
		fmt.Println("Too many students in this room")
		t.Fail()
	} else {
		fmt.Println("TEST PASSED!")
	}

}
