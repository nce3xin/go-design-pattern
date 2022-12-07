package observertemplate

import "testing"

func TestDemo(t *testing.T) {
	subject := &ConcreteSubject{}
	subject.RegisterObserver(&ConcreteObserverOne{})
	subject.RegisterObserver(&ConcreteObserverTwo{})
	subject.NotifyObservers("hello")
}
