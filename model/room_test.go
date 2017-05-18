package model

import (
	"fmt"
	"testing"
)

func TestCheckIfRoomIsOverOccpupied(t *testing.T) {
	room := Room{Name: "G3.076", Occupation: 20, MaxCapacity: 30, Size: 20}

	if room.Occupation > room.MaxCapacity {
		fmt.Println("Name: ", room.Name, "Occupation: ", room.Occupation, "Max Capacity: ", room.MaxCapacity)
		fmt.Println("Too many students in this room")
		t.Fail()
	}

}
