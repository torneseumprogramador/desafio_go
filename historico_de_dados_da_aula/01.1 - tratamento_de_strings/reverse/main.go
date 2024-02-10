package main

import (
	"fmt"
)

func main() {
	str := Reverse("Ola, mundo!")
	fmt.Println(str) // !odnum ,alO
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
