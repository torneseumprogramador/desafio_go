package main

import (
	"fmt"
)

func Soma(x int, y int) int {
	return x + y
}

func main() {
	var r int

	r = Soma(1, 2)

	fmt.Println("Resultado:", r)
}
