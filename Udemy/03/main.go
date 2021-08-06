package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type stringfy interface {
	ToString() string
}

func main() {
	vs := []stringfy{
		&Person{Name: "taro", Age: 21},
		&Car{Number: "4214", Model: "Nissan"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
