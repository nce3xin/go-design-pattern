package example3

// 替换依赖的外部系统

// IA 外部系统A
type IA interface {
	fa()
}

type A struct {
}

func (a A) fa() {
	// ...
}

// Demo 在我们的项目中，外部系统A的使用示例
type Demo struct {
	a IA
}

func NewDemo(a IA) *Demo {
	return &Demo{a: a}
}

func RunViaA() {
	NewDemo(A{})
}

// BAdaptor 实现 IA 接口
type BAdaptor struct {
	b *B
}

func (B BAdaptor) fa() {
	// ...
	B.fa()
}

type B struct {
}

/*
RunViaB 借助BAdaptor，Demo的代码中，调用IA接口的地方都无需改动
   只需要将BAdaptor如下注入到Demo即可
*/
func RunViaB() {
	NewDemo(BAdaptor{})
}
