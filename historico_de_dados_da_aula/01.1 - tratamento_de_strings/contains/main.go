package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Olá, mundo!"
	contains := strings.Contains(str, "mundo")
	fmt.Println(contains)
}
