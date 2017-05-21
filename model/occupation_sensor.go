package model

type OccupationSensor struct {
	InSensorDatapoints  []Datapoint
	OutSensorDatapoints []Datapoint
	TotalEntrances      int
	TotalExits          int
	CurrentOccupants    int
}

func (o *OccupationSensor) CalculateExits() int {

	triggeredOut := false
	// triggeredIn := false

	exits := 0
	for index := range o.OutSensorDatapoints {
		if index == 0 {
			continue
		}

		//Hier moet nog een marge bijkomen van een bepaalde lengte omdat de standaard
		//meting van de deurpost niet altijd gelijk is
		if o.OutSensorDatapoints[index].Value != o.OutSensorDatapoints[index-1].Value {
			triggeredOut = true
		}
		if o.InSensorDatapoints[index].Value != o.InSensorDatapoints[index-1].Value {
			// triggeredIn = true
		}

		if triggeredOut {
			triggeredOut = false
			// triggeredIn = false
			exits++
		}
		index++
	}

	return exits
}

func (o *OccupationSensor) CalculateEntrances() int {
	triggeredIn := false
	// triggeredOut := false

	entrances := 0
	for index := range o.InSensorDatapoints {
		if index == 0 {
			continue
		}

		if o.InSensorDatapoints[index].Value != o.InSensorDatapoints[index-1].Value {
			triggeredIn = true
		}

		if o.OutSensorDatapoints[index].Value != o.OutSensorDatapoints[index-1].Value {
			// triggeredOut = true
		}

		if triggeredIn {
			triggeredIn = false
			// triggeredOut = false
			entrances++
		}
		index++
	}

	return entrances
}

func (o *OccupationSensor) CalculateCurrentOccupants() {

}
