package model

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func SetOccupationSensor() *OccupationSensor {
	occupationSensor := &OccupationSensor{}

	occupationSensor.InSensorDatapoints = CreateDummyDatapointsIn()
	occupationSensor.OutSensorDatapoints = CreateDummyDatapointsOut()
	occupationSensor.TotalEntrances = 0
	occupationSensor.TotalExits = 0
	occupationSensor.CurrentOccupants = 0

	return occupationSensor

}

func CalculateOccupation(occupationSensor *OccupationSensor) int {
	total_exits := occupationSensor.CalculateExits()
	total_entrances := occupationSensor.CalculateEntrances()

	return (total_entrances - total_exits)
}

func CreateDummyDatapointsIn() []*Datapoint {
	datapointsIn := []*Datapoint{}

	number := 60.0
	index := 1
	for index < 5000 {
		datapoint := &Datapoint{}
		datapoint.ID = bson.NewObjectId()
		datapoint.Key = "distance"
		datapoint.Timestamp = time.Now().UnixNano() / 1000000

		if index >= 1000 && index <= 1050 {
			datapoint.Value = 10.0
		} else {
			datapoint.Value = number
		}

		datapointsIn = append(datapointsIn, datapoint)
		index++
	}

	return datapointsIn
}

func CreateDummyDatapointsOut() []*Datapoint {
	datapointsOut := []*Datapoint{}

	number := 60.0
	index := 1
	for index < 5000 {
		datapoint := &Datapoint{}
		datapoint.ID = bson.NewObjectId()
		datapoint.Key = "distance"

		if index >= 1030 && index <= 1050 {
			datapoint.Value = 10.0
		} else {
			datapoint.Value = number
		}

		datapointsOut = append(datapointsOut, datapoint)
		index++
	}

	return datapointsOut
}

func TestOccupationSensor(t *testing.T) {
	sensor := SetOccupationSensor()
	CalculateOccupation(sensor)
}
