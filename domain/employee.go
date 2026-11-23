package domain

type Employee struct {
	Human
	Salary int
}

func (e *Employee) Clone() *Employee {
	if e == nil {
		return nil
	}

	clone := *e

	return &clone
}

func (e *Employee) LookingFor() {}
