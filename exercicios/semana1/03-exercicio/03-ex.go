/* 3.Inicie um número inteiro e converta-o para ponto flutuante e string, exibindo o valor em cada tipo. */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	inteiro := 42
	fmt.Printf("Inteiro: %d", inteiro)
	fmt.Printf("\nString: %s", strconv.Itoa(inteiro))
	fmt.Printf("\nFlutuante: %f", float64(inteiro))
}
