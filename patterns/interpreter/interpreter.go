package interpreter

type Expression interface {
	Interpret() int
}

type Number struct {
	Val int
}

func (n *Number) Interpret() int {
	return n.Val
}

type Add struct {
	Left  Expression
	Right Expression
}

func (a *Add) Interpret() int {
	return a.Left.Interpret() + a.Right.Interpret()
}

type Multiply struct {
	Left  Expression
	Right Expression
}

func (m *Multiply) Interpret() int {
	return m.Left.Interpret() * m.Right.Interpret()
}
