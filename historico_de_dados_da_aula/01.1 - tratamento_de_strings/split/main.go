package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Olá, mundo! Bem-vindo ao Go!"
	parts := strings.Split(str, " ")

	for _, part := range parts {
		fmt.Println(part)
	}
}
