/*
	Exercício 2: Implemente uma interface Forma que tenha um método Area() float64.

Crie duas structs, Quadrado e Círculo, que implementem essa interface.
*/
package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Quadrado struct {
	lado float64
}

type Circulo struct {
	raio float64
}

func (q Quadrado) Area() float64 {
	return q.lado * q.lado
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func imprimeArea(f Forma) {
	fmt.Println(f.Area())
}

func main() {
	q := Quadrado{4.2}
	imprimeArea(q)
	c := Circulo{3.2}
	imprimeArea(c)
}
