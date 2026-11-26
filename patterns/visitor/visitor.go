package visitor

import (
	"fmt"
	"math"
)

type Shape interface {
	Accept(visitor Visitor)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

type Rectangle struct {
	Width  int
	Length int
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}

type Visitor interface {
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
}

type AreaVisitor struct {
	Total float64
}

func (a *AreaVisitor) VisitCircle(c *Circle) {
	area := math.Pi * c.Radius * c.Radius
	a.Total += area
}

func (a *AreaVisitor) VisitRectangle(r *Rectangle) {
	area := r.Length * r.Width
	a.Total += float64(area)
}

type PrintVisitor struct{}

func (p *PrintVisitor) VisitCircle(c *Circle) {
	fmt.Printf("Circle radius: %f\n", c.Radius)
}

func (p *PrintVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("Rectangle length: %f, width: %f\n", float64(r.Length), float64(r.Width))
}
