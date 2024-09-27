/* Exercício 1: Crie uma função que calcule a média de uma lista de números
passados como argumento. */

package main

import "fmt"

func calculaMedia(notas ...float64) float64 {
	total := 0.0
	for _, nota := range notas {
		total += nota
	}
	return total / float64(len(notas))
}

func main() {
	fmt.Println("Media:", calculaMedia(9.8, 7.2, 7.6, 8.5))
}
