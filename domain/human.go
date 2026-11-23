package domain

type TSex string

const (
	Male   TSex = "Male"
	Female TSex = "Female"
)

type Human struct {
	Name string
	Age  int
	Sex  TSex
}
