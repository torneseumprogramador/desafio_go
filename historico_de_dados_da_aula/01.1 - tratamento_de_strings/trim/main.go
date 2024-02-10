package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   Olá, mundo!   "
	trimmedStr := strings.TrimSpace(str)
	fmt.Println(trimmedStr) // "Olá, mundo!"

	str2 := "   Olá, mundo!   "
	lTrimmedStr := strings.TrimLeft(str2, " ")
	fmt.Println(lTrimmedStr) // "Olá, mundo!   "

	str3 := "   Olá, mundo!   "
	rTrimmedStr := strings.TrimRight(str3, " ")
	fmt.Println(rTrimmedStr) // "   Olá, mundo!"

}
