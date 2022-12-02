package idgenerator

import "testing"

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Errorf("More than one instance! 1. %v 2. %v", ins1, ins2)
	}
}
