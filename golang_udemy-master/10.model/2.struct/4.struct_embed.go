package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	Point Point
	A     int
}

func (p *Point) Set(i int) {
	p.A = i
}

func main() {
	bp := BigPoint{}
	fmt.Println(bp)

	bp.Point.A = 100
	bp.Point.B = "Apple"
	bp.Point.C = 2.8

	fmt.Println(bp)

	bp.A = 1000
	fmt.Println(bp.Point.A)
	fmt.Println(bp.A)

	bp2 := BigPoint{
		Point: Point{
			A: 100,
			B: "Banana",
			C: 3.5,
		},
	}

	fmt.Println(bp2)
	bp2.Point.Set(2000)
	fmt.Println(bp2)
}
