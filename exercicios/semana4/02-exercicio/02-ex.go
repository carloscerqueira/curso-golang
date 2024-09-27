/* Exercício 2: Implemente uma função que retorne o maior valor de um slice de inteiros. */
package main

import "fmt"

func buscaMaior(slice []int) int {
	maior := 0
	for _, numero := range slice {
		if maior < numero {
			maior = numero
		}
	}
	return maior
}

func main() {
	slice := make([]int, 3)
	slice[0] = 6
	slice[1] = 7
	slice[2] = 4
	fmt.Println("O maior número do primeiro slice é:", buscaMaior(slice))
	slice2 := []int{1, 6, 7, 3, 87, 123, 3}
	fmt.Println("O maior número do segundo slice é:", buscaMaior(slice2))
}
