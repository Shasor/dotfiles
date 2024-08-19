package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 5555, 55, 93}
	max := Max(a)
	fmt.Println(max)
}