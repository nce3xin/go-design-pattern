package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewOrganization() IOrganization {
	root := &Department{Id: 0}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{Salary: 10})
		root.AddSub(&Department{SubNodes: []IOrganization{&Employee{Salary: 20}}})
	}
	return root
}

func TestDepartment_CalculateSalary(t *testing.T) {
	org := NewOrganization()
	totalSalary := org.CalculateSalary()
	assert.Equal(t, 300, totalSalary)
}
