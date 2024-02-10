package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Ola mundo"
	replacedStr := strings.Replace(str, "mundo", "Go", -1)
	fmt.Println(replacedStr) // "Olá, Go!"
}
