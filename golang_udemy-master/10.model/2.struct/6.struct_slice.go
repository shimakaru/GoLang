package main

import "fmt"

type Person struct {
	Name string
}

type Persons struct {
	Persons []*Person
}

func main() {
	ps := make([]Person, 5)
	fmt.Println(ps)

	ps[0].Name = "Mike"
	ps[1].Name = "Jon"
	ps[2].Name = "Nina"

	fmt.Println(ps)

	p1 := Person{"Mike"}
	p2 := Person{"Jon"}
	p3 := Person{"Nina"}
	ps2 := Persons{}

	ps2.Persons = append(ps2.Persons, &p1, &p2, &p3)

	for _, p := range ps2.Persons {
		fmt.Println(p)
	}
}
