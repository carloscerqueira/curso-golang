/* Exercício 1: Crie um array de inteiros e faça um programa que calcule
a soma de todos os elementos. */

package main

import "fmt"

func main() {
	arrayInteiros := [10]int{6, 7, 3, 4, 5, 7, 8, 7, 8, 9}
	soma := 0
	for _, numero := range arrayInteiros {
		soma += numero
	}
	fmt.Printf("A soma dos números do array é: %d", soma)
}
