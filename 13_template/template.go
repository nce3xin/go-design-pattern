package template

import "fmt"

type IWorker interface {
	getUp()
	eatBreakfast()
	goToWork()
}

// Worker IWork基类
type Worker struct {
	// struct嵌套interface，其实就是持有了一个类型是interface的成员变量
	// NewWorker()时，传入的参数类型必须是实现了IWorker接口的类型
	IWorker
}

// 模板方法
func (w *Worker) work() {
	w.getUp()
	w.eatBreakfast()
	w.goToWork()
}

// Programmer 具体类：程序员工作流程
type Programmer struct {
	Worker
}

func (p *Programmer) getUp() {
	fmt.Printf("programmer get up\n")
}

func (p *Programmer) eatBreakfast() {
	fmt.Printf("programmer eat breakfast\n")
}

func (p *Programmer) goToWork() {
	fmt.Printf("programmer go to work\n")
}

// Doctor 具体类：医生工作流程
type Doctor struct {
	Worker
}

func (d *Doctor) getUp() {
	fmt.Printf("doctor get up\n")
}

func (d *Doctor) eatBreakfast() {
	fmt.Printf("doctor eat breakfast\n")
}

func (d *Doctor) goToWork() {
	fmt.Printf("doctor go to work\n")
}
