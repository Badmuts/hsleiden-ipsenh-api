package model

type OccupationSensor struct {
	InSensorDatapoints  []*Datapoint
	OutSensorDatapoints []*Datapoint
	TotalEntrances      int
	TotalExits          int
	CurrentOccupants    int
}

func (o *OccupationSensor) CalculateExits() int {

	triggeredOut := false
	triggeredIn := false

	exits := 0
	for index := range o.OutSensorDatapoints {
		if o.OutSensorDatapoints[index].Value != o.OutSensorDatapoints[index-1].Value {
			triggeredOut = true
		}
		if o.InSensorDatapoints[index].Value != o.InSensorDatapoints[index].Value {
			triggeredIn = true
		}

		if triggeredIn && triggeredOut {
			triggeredOut = false
			triggeredIn = false
			exits++
		}
	}

	return exits
}

// func (o *OccupationSensor) CalculateEntrances() int {

// }
