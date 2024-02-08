package main

import (
	"fmt"
)

func main() {
	var idade int = 80
	// fmt.Print("Digite sua idade: ")
	// fmt.Scanln(&idade)

	if idade < 13 {
		fmt.Println("Criança")
	} else if idade >= 13 && idade < 18 {
		fmt.Println("Adolescente")
	} else if idade >= 18 && idade < 65 {
		fmt.Println("Adulto")
	} else {
		fmt.Println("Idoso")
	}

	switch {
	case idade < 13:
		fmt.Println("Criança")
	case idade >= 13 && idade < 18:
		fmt.Println("Adolescente")
	case idade >= 18 && idade < 65:
		fmt.Println("Adulto")
	default:
		fmt.Println("Idoso")
	}

	if idade == 13 {
		fmt.Println("Criança")
	} else if idade == 14 {
		fmt.Println("Adolescente")
	} else if idade == 18 {
		fmt.Println("Adulto")
	} else {
		fmt.Println("Outros")
	}

	switch idade {
	case 13:
		fmt.Println("Criança")
	case 14:
		fmt.Println("Adolescente")
	case 18:
		fmt.Println("Adulto")
	default:
		fmt.Println("Outros")
	}
}
