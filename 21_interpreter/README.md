# 解释器模式

解释器模式为某个语言定义它的语法表示，并定义一个解释器来处理这个语法。

## 实现关键

将语法解析的工作拆分到各个小类中，以此来避免大而全的解析类。

将语法规则拆分成一些小的独立的单元，然后对每个单元进行解析，最终合并为对整个语法规则的解析。