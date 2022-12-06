package adapter

import "testing"

func TestAdapter_f1(t *testing.T) {
	a := Adapter{adaptee: &Adaptee{}}
	a.f1()
	a.f2()
	a.fc()
}
