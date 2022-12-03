package prototype

import "encoding/json"

type Employee struct {
	Name    string
	Addr    *Address
	Company *Company
	Gender  string
}

type Address struct {
	Province string
	City     string
	Street   string
}

type Company struct {
	Name string
	Id   string
}

func (e *Employee) DeepClone() *Employee {
	// 通过序列化反序列化的方法，实现深拷贝
	var newE Employee
	bytes, _ := json.Marshal(e)
	json.Unmarshal(bytes, &newE)
	return &newE
}
