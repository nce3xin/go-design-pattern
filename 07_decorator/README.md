# 装饰器模式

装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。组合优于继承。

装饰器模式和代理模式很像，代码结构基本都是一致的。

```go
type IA interface {
	f()
}

type A struct {
}

func (a *A) f() {
	// TODO
}

type ADecorator struct {
	A IA
}

func (a *ADecorator) f() {
    // 功能增强代码
    a.A.f()
    // 功能增强代码
}
```



