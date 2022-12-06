package facade

type Fridge struct {
}

func (f *Fridge) turnOff() {

}

type Television struct {
}

func (t *Television) turnOff() {

}

// Facade 门面类
type Facade struct {
	f *Fridge
	t *Television
}

func NewFacade() *Facade {
	return &Facade{
		f: &Fridge{},
		t: &Television{},
	}
}

func (f *Facade) turnOff() {
	f.f.turnOff()
	f.t.turnOff()
	return
}
