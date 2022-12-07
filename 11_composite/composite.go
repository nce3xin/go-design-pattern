package composite

type IOrganization interface {
	CalculateSalary() int
}

type Employee struct {
	Id     int
	Salary int
}

func (e *Employee) CalculateSalary() int {
	return e.Salary
}

type Department struct {
	Id       int
	Salary   int
	SubNodes []IOrganization
}

func (d *Department) CalculateSalary() int {
	totalSalary := 0
	for _, node := range d.SubNodes {
		totalSalary += node.CalculateSalary()
	}
	d.Salary = totalSalary
	return totalSalary
}

func (d *Department) AddSub(org IOrganization) {
	d.SubNodes = append(d.SubNodes, org)
}
