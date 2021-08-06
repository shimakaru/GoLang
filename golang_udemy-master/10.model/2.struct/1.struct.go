package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

func main() {
	var p Point
	fmt.Println(p)

	p2 := Point{A: 1, B: "Go", C: 1.2}
	fmt.Println(p2)

	p2.A = 10
	fmt.Println(p2)

	p3 := Point{1, "python", 3.14}
	fmt.Println(p3)

	p4 := Point{A: 2}
	fmt.Println(p4)
}
