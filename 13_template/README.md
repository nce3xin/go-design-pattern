# 模板方法模式

模板方法模式在一个方法中定义一个业务逻辑骨架（就是模板），并将某些步骤推迟到子类中实现。可以让子类在不改变业务逻辑整体结构的情况下，重新定义业务逻辑中的某些步骤。包含业务逻辑骨架（模板）的方法就是“模板方法”。这也是模板方法模式名字的由来。

由于go语言不支持继承，这里给出一个java经典版的实现。templateMethod()函数定义为final，是为了避免子类重写它。method1() 和 method2() 定义为abstract，是为了强迫子类去实现。不过，这些都不是必须的，在实际的项目开发中，模板模式的代码实现比较灵活。

```java
public abstract class AbstractClass{
    public final void templateMethod(){
        // ...
        method1();
        // ...
        method2();
        // ...
    }
    
    protected abstract void method1();
    protected abstract void method2();
}

public class ConcreteClass1 extends AbstractClass{
    @override
    protected void method1(){
        // ...
    }
    
    @override
    protected void method2(){
        // ...
    }
}

public class ConcreteClass2 extends AbstractClass{
    @override
    protected void method1(){
        // ...
    }
    
    @override
    protected void method2(){
        // ...
    }
}

AbstractClass demo = ConcreteClass1();
demo.templateMethod();
```

## 模板方法模式的作用

- 复用。
- 扩展。

### 1. 复用

模板方法模式把一个业务逻辑中，不变的流程抽象到父类的模板方法 templateMethod() 中，将可变的部分 method1() 和 method2() 留给子类 ConcreteClass1 和 ConcreteClass2 来实现。所有的子类都可以复用父类中模板方法 templateMethod() 定义的流程代码。

### 2. 扩展

这里的扩展，不是指代码的扩展性，而是指框架的扩展性。模板方法模式常用在框架的开发中，让框架用户可以在不修改框架源码的情况下，定制化框架的功能。

## 回调

还有一个技术概念，也能起到和模板模式相同的作用，就是回调（Callback）。

回调是一种双向调用关系。A类事先注册某个函数F到B类，A类在调用B类的P函数的时候，B类反过来调用A类注册给它的F函数。这里的 F 函数就是回调函数。A调用B，B反过来又调用A，这种调用机制就叫作回调。
