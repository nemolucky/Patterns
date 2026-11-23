package factory

import "github.com/nemolucky/Patterns/domain"

type EmployeeFactory struct{}

func (f *EmployeeFactory) NewEmployee(id int, name, sex string, age, salary int) *domain.Employee {
	return &domain.Employee{
		Name:   name,
		Sex:    domain.TSex(sex),
		Age:    age,
		Salary: salary,
	}
}
