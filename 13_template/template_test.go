package template

import "testing"

func TestTemplate(t *testing.T) {
	worker1 := &Worker{&Programmer{}}
	worker1.work()

	worker2 := &Worker{&Doctor{}}
	worker2.work()
}
