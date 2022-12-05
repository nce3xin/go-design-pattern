package decorator

import "testing"

func TestInputStreamRead(t *testing.T) {
	// 装饰器类和原始类实现相同的接口，这样可以对原始类嵌套多个装饰器类
	in := &FileInputStream{}
	bin := BufferedInputStream{in: in}
	din := DataInputStream{in: bin}
	din.ReadInt()
}
