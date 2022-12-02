# 建造者模式

在go中，建造者模式可以用option来实现。

## 为什么要用建造者模式

在我们编程中，我们会经常性的需要对一个对象（或是业务实体）进行相关的配置。比如下面这个业务实体：

```go
type Server struct {
    Addr     string
    Port     int
    Protocol string
    Timeout  time.Duration
    MaxConns int
    TLS      *tls.Config
}
```

Server有些属性是必填的，比如ip和端口号。Protocol，Timeout，MaxConns是不能为空，但有默认值的。TLS需要配置相关的证书，可以为空。

所以，针对于上述这样的配置，我们需要有多种不同的构造函数，创建不同配置的Server。

为了高效解决这个问题，出现了建造者模式。

## Go实现方式

函数式编程 functional option。

## 问题

### 1. 如果想要验证创建的对象是否合法，比如字段之间的依赖关系等等，怎么办？

可以在调用完所有的 Option 后，对创建好的对象进行验证，就像是在 build 这一步中做一样。





