package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "(11) 11111-1111"
	re := regexp.MustCompile(`^\(\d{2}\) \d{5}-\d{4}$`)
	match := re.MatchString(str)
	fmt.Println(match)
}
