package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Meu número é (11) 02999-9799. dsdssd ds sdd sdd ssd sd"
	re := regexp.MustCompile(`\(\d{2}\) \d{5}-\d{4}`)
	match := re.FindString(str)
	fmt.Println(match)
}
