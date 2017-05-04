package main

import (
	"fmt"
)

type Datapoint struct {
	sensor    Sensor
	key       int
	value     float32
	timestamp int //dit moet nog een timestamp worden
}

func main() {
	fmt.Println("Datapoint class")
}
