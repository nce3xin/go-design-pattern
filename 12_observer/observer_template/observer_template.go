package observertemplate

// ISubject 被观察者
type ISubject interface {
	RegisterObserver(o IObserver)
	RemoveObserver(o IObserver)
	NotifyObservers(msg string)
}

type IObserver interface {
	Update(msg string)
}

type ConcreteSubject struct {
	observers []IObserver
}

func (c *ConcreteSubject) RegisterObserver(o IObserver) {
	c.observers = append(c.observers, o)
}

func (c *ConcreteSubject) RemoveObserver(o IObserver) {
	removeIdx := -1
	for i, observer := range c.observers {
		if observer == o {
			removeIdx = i
			break
		}
	}
	c.observers = append(c.observers[:removeIdx], c.observers[removeIdx+1:]...)
}

func (c *ConcreteSubject) NotifyObservers(msg string) {
	for _, observer := range c.observers {
		observer.Update(msg)
	}
}

type ConcreteObserverOne struct {
}

func (c ConcreteObserverOne) Update(msg string) {
	//TODO implement me
}

type ConcreteObserverTwo struct {
}

func (c ConcreteObserverTwo) Update(msg string) {
	//TODO implement me
}




