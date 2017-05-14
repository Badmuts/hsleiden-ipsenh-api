package model

type SensorType interface {
	current() float32
	average() float32
}
