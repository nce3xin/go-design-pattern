package adapter

type ITarget interface {
	f1()
	f2()
	fc()
}

type Adaptee struct {
}

func (a *Adaptee) fa() {
	// ...
}

func (a *Adaptee) fb() {
	// ...
}

func (a *Adaptee) fc() {
	// ...
}

type Adapter struct {
	adaptee *Adaptee
}

func (a Adapter) f1() {
	a.adaptee.fa()
}

func (a Adapter) f2() {
	a.adaptee.fb()
}

func (a Adapter) fc() {
	a.adaptee.fc()
}
