/* Exercício 4: Escreva uma função anônima que filtre números
pares de um slice e retorne um novo slice contendo apenas os números pares. */

package main

import "fmt"

func main() {
	filtraPares := func(slice []int) []int {
		sliceResult := make([]int, 0)
		for _, numero := range slice {
			if numero%2 == 0 {
				sliceResult = append(sliceResult, numero)
			}
		}
		return sliceResult
	}

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}
	fmt.Println("Slice final: ", filtraPares(slice))
}
