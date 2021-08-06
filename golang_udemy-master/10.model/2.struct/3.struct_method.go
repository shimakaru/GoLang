package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

func Set(p *Point, i int) {
	p.A = i
}

func (p *Point) Set(i int) {
	p.A = i
}

func (p Point) Set2(i int) {
	p.A = i
}

func main() {
	p1 := &Point{A: 1}
	Set(p1, 2)
	fmt.Println(p1)

	p1.Set(3)
	fmt.Println(p1)

	p2 := Point{}
	p2.Set2(100)
	fmt.Println(p2)

	p2.Set(200)
	fmt.Println(p2)
}
