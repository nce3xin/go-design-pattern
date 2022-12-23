package mediator

import "fmt"

type IMediator interface {
	CanComeInStation(train ITrain) bool
	NotifyAboutDeparture()
}

type ITrain interface {
	Arrive()
	Depart()
}

// PassengerTrain 载客火车
type PassengerTrain struct {
	mediator IMediator
}

func (p *PassengerTrain) Arrive() {
	if !p.mediator.CanComeInStation(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.mediator.NotifyAboutDeparture()
}

// FreightTrain 载货火车
type FreightTrain struct {
	mediator IMediator
}

func (f *FreightTrain) Arrive() {
	if !f.mediator.CanComeInStation(f) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (f FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	f.mediator.NotifyAboutDeparture()
}

type StationManager struct {
	trainQueue     []ITrain
	isPlatformFree bool
}

func (s *StationManager) CanComeInStation(train ITrain) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, train)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		train := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		train.Arrive()
	}
}
