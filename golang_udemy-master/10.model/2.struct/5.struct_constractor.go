package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

func NewPoint(a int, b string, c float64) *Point {
	return &Point{a, b, c}
}

func main() {
	p1 := Point{1, "A", 1.1}
	p2 := NewPoint(2, "B", 2.2)
	fmt.Println(p1, p2)
}
