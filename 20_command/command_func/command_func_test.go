package commandfunc

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	// 模拟客户端请求
	eventChan := make(chan string)
	go func() {
		events := []string{"start", "save", "exit"}
		for _, event := range events {
			eventChan <- event
		}
	}()
	defer close(eventChan)

	// 使用命令队列缓存命令
	commands := make(chan Command, 100)
	defer close(commands)

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}
			var command Command
			switch event {
			case "start":
				command = StartCommand()
			case "save":
				command = SaveCommand()
			case "exit":
				command = ExitCommand()
			}
			commands <- command
		}
	}()

	for {
		select {
		case command := <-commands:
			command()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
