package main

import "fmt"

func main() {

	// var i int = 42

	// var p *int // ponteiro definido
	// p = &i     // passagem da referencia de i

	// var t *int // ponteiro definido
	// t = &i     // passagem da referencia de i

	// var u *int // ponteiro definido
	// u = &i     // passagem da referencia de i

	// i = 20

	// *p = 30

	// fmt.Printf("Valor de i (%d) o endereço de memória (%p)\n", i, &p)
	// fmt.Printf("Valor de p (%d) o endereço de memória (%p)\n", *p, p)
	// fmt.Printf("Valor de t (%d) o endereço de memória (%p)\n", *t, t)
	// fmt.Printf("Valor de u (%d) o endereço de memória (%p)\n", *u, u)

	// fmt.Println("======== valor por copia ============")

	fmt.Println("======== valor por referencia ============")

	t := 42
	w := &t

	fmt.Printf("Valor de t (%d) o endereço de memória (%p)\n", t, &t)
	fmt.Printf("Valor de w (%d) o endereço de memória (%p)\n", *w, w)

	fmt.Println("======== valor por copia ============")

	x := 42
	z := x

	fmt.Printf("Valor de x (%d) o endereço de memória (%p)\n", x, &x)
	fmt.Printf("Valor de z (%d) o endereço de memória (%p)\n", z, &z)

}
