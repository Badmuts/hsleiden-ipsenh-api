package model

import (
	"fmt"
	"testing"
)

func TestCheckIfRoomIsOverOccpupied(t *testing.T) {
	room := Room{Name: "G3.076", Occupation: 20, maxCapacity: 30, Size: 20}

	if room.Occupation > room.maxCapacity {
		fmt.Println("Name: ", room.Name, "Occupation: ", room.Occupation, "Max Capacity: ", room.maxCapacity)
		fmt.Println("Too many students in this room")
		t.Fail()
	}

}
