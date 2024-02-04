package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex) Soma() float64 {
	v.X += 1
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Objeto inspecionado( %+v ), retorno da soma: %v\n", v, v.Soma())
	v.Scale(5)
	fmt.Printf("Objeto inspecionado( %+v ), retorno da soma: %v\n", v, v.Soma())
}
