package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v Vertex) RetornoDaSoma(f float64) float64 {
	return v.X + v.Y
}

func (v *Vertex) SomaAlterarInstancia(f float64) {
	v.X += f
	v.Y += f
}

func SomaAlterandoStruct(v *Vertex, f float64) {
	v.X += f
	v.Y += f
}

func main() {
	v := Vertex{3, 4} // verificar se existe a idéia de construtor

	fmt.Printf("x: %.2f, y: %.2f\n", v.X, v.Y)

	r := v.RetornoDaSoma(1)
	fmt.Printf("x: %.2f, y: %.2f = %.2f\n", v.X, v.Y, r)

	v.SomaAlterarInstancia(1)
	fmt.Printf("x: %.2f, y: %.2f\n", v.X, v.Y)

	SomaAlterandoStruct(&v, 10)
	fmt.Printf("x: %.2f, y: %.2f\n", v.X, v.Y)
}
