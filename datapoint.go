package main

import (
	"fmt"
)

type Datapoint struct {
	sensor    Sensor
	key       int
	value     float32
	timestamp time
}

func main() {
	fmt.Println("Datapoint class")
}
