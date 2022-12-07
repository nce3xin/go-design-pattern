package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type IBus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{}) error
}

type IBusController interface {
	WaitAsync()
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
	wg       sync.WaitGroup
}

func (a *AsyncEventBus) WaitAsync() {
	a.wg.Wait()
}

func (a *AsyncEventBus) Subscribe(topic string, handler interface{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	v := reflect.ValueOf(handler)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}
	_, ok := a.handlers[topic]
	if !ok {
		a.handlers[topic] = make([]reflect.Value, 0)
	}
	a.handlers[topic] = append(a.handlers[topic], v)
	return nil
}

func (a *AsyncEventBus) Publish(topic string, args ...interface{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	handlers, ok := a.handlers[topic]

	if !ok {
		return fmt.Errorf("cannot find handlers under %v topic", topic)
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		v := reflect.ValueOf(arg)
		params[i] = v
	}

	for _, handler := range handlers {
		a.wg.Add(1)
		go a.Call(handler, params)
	}
	return nil
}

func (a *AsyncEventBus) Call(handler reflect.Value, params []reflect.Value) {
	defer a.wg.Done()
	handler.Call(params)
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: make(map[string][]reflect.Value, 0),
		lock:     sync.Mutex{},
		wg:       sync.WaitGroup{},
	}
}
