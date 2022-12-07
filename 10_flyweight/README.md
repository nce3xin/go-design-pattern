# 享元模式 Flyweight

所谓“享元”，顾名思义就是被共享的单元。享元模式的意图是复用对象，节省内存，前提是享元对象是不可变对象。

具体来讲，当一个系统中存在大量重复对象的时候，如果这些重复的对象是不可变对象，我们就可以利用享元模式将对象设计成享元，在内存中只保存一份实例，供多处代码引用。这样可以减少内存中对象的数量，起到节省内存的目的。实际上，不仅仅相同对象可以设计成享元，对于相似对象，我们也可以将这些对象中相同的部分（字段）提取出来，设计成享元，让这些大量相似对象引用这些享元。

这里解释一下，定义中的“不可变对象”，指的是，一旦通过构造函数初始化完成之后，它的状态（对象的成员变量或者属性）就不会再被修改了。所以，不可变对象不能暴露任何 set() 等修改内部状态的方法。之所以要求享元是不可变对象，那是因为它会被多处代码共享使用，避免一处代码对享元进行了修改，影响到其他使用它的代码。

下面通过一个简单的例子解释一下享元模式。

## 棋牌游戏的例子

假设我们在开发一个棋牌游戏（比如象棋）。一个游戏厅中有成千上万个“房间”。每个房间对应一个棋局。棋局要保存每个棋子的数据，比如：棋子类型（车、马、炮等），棋子颜色（红、黑）、棋子在棋局中的位置。

为了记录每个房间当前的棋局情况，我们需要给每个房间都创建一个ChessBoard棋局对象。因为大厅有成千上万的棋局，那保存这么多棋局对象就会消耗大量的内存，有没有什么办法来节省内存呢？

这个时候，享元模式就可以派上用场了。每个棋子对象的id、text、color都是相同的，唯独position X和position Y不同。实际上，我们可以将棋子的id、text、color属性拆分出来，设计成独立的类，并且作为享元供多个棋盘复用。这样，棋盘只需要记录每个棋子的位置就可以了。具体实现见代码。

在使用享元模式之前，记录1万个棋局，我们要创建32万（1w*32）个棋子对象。利用享元模式，我们只需要创建30个享元对象供所有棋局共享使用，大大节省了内存。

## 享元模式的实现

很简单，通过一个map或者list来缓存已经创建好的享元对象，以达到复用目的。

## 总结

在 Java Integer的实现中，-128到127之间的整型对象会被事先创建好，缓存在IntegerCache类中，当我们使用自动装箱或者valueOf()来创建这个数值区间的整型对象时，会复用IntegerCache类事先创建好的对象。这里的IntegerCache类就是享元工厂类，事先创建好的整型对象就是享元对象。

在 Java string类的实现中，JVM开辟一块存储区专门存储字符串常量，这块存储区叫字符串常量池，类似于Integer中的IntegerCache。不过，跟IntegerCache不同的是，它并非事先创建好需要共享的对象，而是在程序的运行期间，根据需要来创建和缓存字符串常量。

除此之外，再补充一点，实际上，享元模式对JVM的垃圾回收并不友好。因为享元工厂类一直保存了对享元对象的引用，这就导致享元对象在没有任何代码使用的情况下，也并不会被JVM垃圾回收机制自动回收掉。因此，在某些情况下，如果对象的生命周期很短，也不会被密集使用，利用享元模式反倒可能会浪费更多的内存。所以，除非经过线上验证，利用享元模式真的可以大大节省内存，否则，就不要过度使用这个模式，为了一点点内存的节省而引入一个复杂的设计模式，得不偿失啊。