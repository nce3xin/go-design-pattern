package example1

// 封装有缺陷的接口设计

// CD 这个类来自外部sdk，我们无权修改
type CD struct {
}

func (cd *CD) uglyNamingFunction1() {

}

func (cd *CD) tooManyParamsFunction2(paramA, paramB int) {

}

func (cd *CD) lowPerformanceFunction3() {

}

// ITarget 使用适配器模式进行重构
type ITarget interface {
	function1()
	function2(paramsWrapper ParamsWrapper)
	function3()
}

type ParamsWrapper struct {
}

func (p *ParamsWrapper) getParamA() int {
	return 0
}

func (p *ParamsWrapper) getParamB() int {
	return 0
}

// CDAdaptor 适配器类
type CDAdaptor struct {
	cd *CD
}

func (C CDAdaptor) function1() {
	C.cd.uglyNamingFunction1()
}

func (C CDAdaptor) function2(paramsWrapper ParamsWrapper) {
	C.cd.tooManyParamsFunction2(paramsWrapper.getParamA(), paramsWrapper.getParamB())
}

func (C CDAdaptor) function3() {
	// ...reimplement it...
}
