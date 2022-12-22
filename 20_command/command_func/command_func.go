package commandfunc

import "fmt"

// 这是示例二，采用将直接返回一个函数，不用对象
// 假设现在有一个游戏服务，我们正在实现一个游戏后端
// 使用一个 goroutine 不断接收来自客户端请求的命令，并且将它放置到一个队列当中
// 然后我们在另外一个 goroutine 中来执行它

type Command func() error

func StartCommand() Command {
	return func() error {
		fmt.Printf("game starting\n")
		return nil
	}
}

func SaveCommand() Command {
	return func() error {
		fmt.Printf("game saving\n")
		return nil
	}
}

func ExitCommand() Command {
	return func() error {
		fmt.Printf("game exiting\n")
		return nil
	}
}
