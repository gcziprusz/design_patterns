package main

import "fmt"

type object interface {
	accept(*visitor)
	name() string
}
type visitor interface {
	visitCircle(*circle)
	visitSquare(*square)
	visitTriangle(*triangle)
}

type circle struct{}

func (c *circle) accept(v visitor) {
	v.visitCircle(c)
}
func (c *circle) name() string {
	return "circle"
}

type triangle struct{}

func (t *triangle) accept(v visitor) {
	v.visitTriangle(t)
}
func (t *triangle) name() string {
	return "triangle"
}

type square struct{}

func (s *square) accept(v visitor) {
	v.visitSquare(s)
}
func (s *square) name() string {
	return "square"
}

type areaCalculator struct{}

func (a *areaCalculator) visitCircle(c *circle) {
	fmt.Printf("Calculate Area of %s\n", c.name())
}
func (a *areaCalculator) visitSquare(s *square) {
	fmt.Printf("Calculate Area of %s\n", s.name())
}
func (a *areaCalculator) visitTriangle(t *triangle) {
	fmt.Printf("Calculate Area of a %s\n", t.name())
}

type midPointCalculator struct{}

func (m *midPointCalculator) visitCircle(c *circle) {
	fmt.Printf("Calculate MidPoint of %s\n", c.name())
}
func (m *midPointCalculator) visitSquare(s *square) {
	fmt.Printf("Calculate MidPoint of %s\n", s.name())
}
func (m *midPointCalculator) visitTriangle(t *triangle) {
	fmt.Printf("Calculate MidPoint of a %s\n", t.name())
}
func main() {
	var ci = &circle{}
	var sq = &square{}
	var tr = &triangle{}

	var ac = &areaCalculator{}

	ci.accept(ac)
	sq.accept(ac)
	tr.accept(ac)

	fmt.Println("------------------------")
	var mi = &midPointCalculator{}

	ci.accept(mi)
	sq.accept(mi)
	tr.accept(mi)
}
