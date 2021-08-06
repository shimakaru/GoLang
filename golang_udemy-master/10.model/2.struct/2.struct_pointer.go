package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

func Update(p Point) {
	p.A = 100
	p.B = "Update"
	p.C = 2.14
}

func Update2(p *Point) {
	p.A = 100
	p.B = "Update"
	p.C = 2.14
}

func main() {
	var p Point
	Update(p)
	fmt.Println(p)

	p2 := &Point{}
	Update2(p2)
	fmt.Println(p2)

	Update2(&p)
	fmt.Println(p)

	p3 := new(Point)
	Update2(p3)
	fmt.Println(p3)
}
