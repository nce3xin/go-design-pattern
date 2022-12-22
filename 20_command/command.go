package command

import "fmt"

// 假设现在有一个游戏服务，我们正在实现一个游戏后端
// 使用一个 goroutine 不断接收来自客户端请求的命令，并且将它放置到一个队列当中
// 然后我们在另外一个 goroutine 中来执行它

type ICommand interface {
	Execute() error
}

type StartCommand struct {
}

func (s *StartCommand) Execute() error {
	fmt.Printf("game starting\n")
	return nil
}

type SaveCommand struct {
}

func (s *SaveCommand) Execute() error {
	fmt.Printf("game saving\n")
	return nil
}

type ExitCommand struct {
}

func (e *ExitCommand) Execute() error {
	fmt.Printf("game exiting\n")
	return nil
}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func NewSaveCommand() *SaveCommand {
	return &SaveCommand{}
}

func NewExitCommand() *ExitCommand {
	return &ExitCommand{}
}
