package main

import "fmt"

type Square struct{
	sideLength float64
}
type Triangle struct{
	base float64
	height float64
}

type Shape interface {
	getArea() float64
}

func main() {
	s := Square{5}
	t := Triangle{5, 10}
	printArea(s)
	printArea(t)

}


func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
