package main

import (
	"06/alib"
	"fmt"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	ii := 1
	fmt.Println(IsOne(ii))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
