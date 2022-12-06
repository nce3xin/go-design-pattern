package facade

import "testing"

func TestFacadeTurnoff(t *testing.T) {
	f := NewFacade()
	f.turnOff()
}
