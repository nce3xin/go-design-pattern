package command

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
	commands := make(chan ICommand, 100)
	defer close(commands)

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}
			var command ICommand
			switch event {
			case "start":
				command = NewStartCommand()
			case "save":
				command = NewSaveCommand()
			case "exit":
				command = NewExitCommand()
			}
			commands <- command
		}
	}()

	for {
		select {
		case command := <-commands:
			command.Execute()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
