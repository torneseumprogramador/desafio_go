package main

var globalVar int = 10 // memória static
const PI = 3.14        // memória static

func main() {
	println(PI) // imprimindo constante

	globalVar = 2323 // alterando variável global

	numero := 5 // memória stack
	println(numero)

	array := [2]int{1, 3} // memória stack
	println(array[0])

	// memória heap (pois é um dado muito grande para ficar na stack)
	textoDeUmLivro := "32823083270232309723098230932 ...322323 32,, 1 GB de string" // memória stack
	println(textoDeUmLivro)
}
