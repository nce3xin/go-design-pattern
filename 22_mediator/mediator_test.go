package mediator

import "testing"

func TestDemo(t *testing.T) {
	sm := &StationManager{
		trainQueue:     make([]ITrain, 0),
		isPlatformFree: true,
	}
	pt := &PassengerTrain{mediator: sm}
	ft := &FreightTrain{mediator: sm}
	pt.Arrive()
	ft.Arrive()
	pt.Depart()
}
