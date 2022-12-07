# 观察者模式

在对象之间定义一个一对多的依赖，当一个对象状态改变的时候，所有依赖的对象都会自动收到通知。

一般来说，被观察的对象叫被观察者Observable，观察者叫Observer。不过在实际的项目开发中，这两种对象的称呼是比较灵活的，有各种不同的叫法，比如：

- Subject - Observer
- Publisher - Subscriber
- Producer - Consumer
- Event Emitter - Event Listener
- Dispatcher - Listener

不管怎么称呼，只要应用场景符合刚刚给出的定义，都可以看作观察者模式。

实际上，观察者模式是一个比较抽象的模式，根据不同的应用场景和需求，有完全不同的实现方式。先来看一种最经典的，模板化的实现方式：

```go
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
```

实际上，上面的代码算是观察者模式的“模板代码”，只能反映大体的设计思路。在真实的软件开发中，并不需要照搬上面的模板代码。观察者模式的实现方法各式各样，函数、类的命名会根据业务场景的不同有很大的差别，比如register函数还可以叫做attach，remove函数还可以叫做detach等等。不过，万变不离其宗，设计思路都是差不多的。

## 例子

一个投资理财平台，用户注册成功之后，会给用户发放优惠券，并发送一封欢迎邮件。见代码。

## 基于不同应用场景的不同实现方式

观察者模式的应用场景非常广泛，小到代码层面的解耦，大到架构层面的系统解耦，再或者一些产品的设计思路，都有这种模式的影子，比如，邮件订阅、RSS Feeds，本质上都是观察者模式。

有同步阻塞、异步非阻塞、进程内、跨进程的实现方式。

刚刚的实现方式，是同步阻塞的。观察者和被观察者代码在同一个线程内执行，被观察者一直阻塞，直到所有的观察者代码都执行完成之后，才执行后续的代码。

如何实现异步非阻塞呢？简单一点的做法，是每个Update()函数中，创建一个新的线程执行代码。不过，还有更优雅的实现，就是EventBus。

刚刚讲到的2个场景，不管是同步阻塞还是异步非阻塞，都是进程内的实现方式。如果要跨进程呢？最好的方式就是基于消息队列（Message Queue）来实现。

