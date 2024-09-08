/* 8. Crie um programa que declare uma variável inteira e um ponteiro para essa variável. Use o ponteiro para alterar o valor
da variável original e imprima o valor antes e depois da alteração. */

package main

import "fmt"

func main() {
	inteiro := 42
	var ponteiro *int = &inteiro

	fmt.Printf("Valor inicial: %d", inteiro)
	*ponteiro = 40
	fmt.Printf("\nValor final: %d", inteiro)
}
