# Go设计模式

Go设计模式实现。

## 目录

### 创建型
- [简单工厂模式](https://github.com/nce3xin/go-design-pattern/tree/main/00_simple_factory)
- [工厂方法模式](https://github.com/nce3xin/go-design-pattern/tree/main/01_factory_method)
- [单例模式](https://github.com/nce3xin/go-design-pattern/tree/main/02_singleton)
- [建造者模式](https://github.com/nce3xin/go-design-pattern/tree/main/03_builder)
- [原型模式](https://github.com/nce3xin/go-design-pattern/tree/main/04_prototype)

创建型模式主要解决对象的创建问题，封装复杂的创建过程，解耦对象的创建代码和使用代码。

其中，单例模式用来创建全局唯一的对象。工厂模式用来创建不同但是相关类型的对象（继承同一父类或者接口的一组子类）。由给定的参数来决定创建哪种类型的对象。建造者模式是用来创建复杂对象，可以通过设置不同的可选参数，“定制化” 创建不同的对象。原型模式针对创建成本比较大的对象，利用对已有对象进行复制的方式进行创建，以达到节省创建时间的目的。

## 闲言碎语

### Goland环境变量配置

```
GOPROXY=https://mirrors.aliyun.com/goproxy/,direct;GO111MODULE=on
```

### go.mod生成方法

```
go mod init <name>
```

