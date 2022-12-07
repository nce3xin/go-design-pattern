package eventbus

import (
	"fmt"
	"testing"
)

func TestAsyncEventBus_Publish(t *testing.T) {
	asyncBus := NewAsyncEventBus()
	asyncBus.Subscribe("topic:1", handlerOne)
	asyncBus.Subscribe("topic:1", handlerTwo)
	asyncBus.Publish("topic:1", "hello", "world")

	fmt.Println("start: do some stuff while waiting for a result")
	fmt.Println("end: do some stuff while waiting for a result")

	asyncBus.WaitAsync()

	fmt.Println("do some stuff after waiting for result")
}

func handlerOne(arg1, arg2 string) {
	fmt.Printf("handlerOne: %s, %s\n", arg1, arg2)
}

func handlerTwo(arg1, arg2 string) {
	fmt.Printf("handlerTwo: %s, %s\n", arg1, arg2)
}
