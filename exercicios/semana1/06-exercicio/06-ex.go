/* 6. Crie um programa que inicializa dois números inteiros e exiba os resultados das comparações usando os operadores
relacionais (==, !=, <, >, <=, >=). */

package main

import "fmt"

func main() {
	a, b := 4, 5
	fmt.Printf("a = %d\nb = %d", a, b)
	fmt.Printf("\na == b: %t", a == b)
	fmt.Printf("\na != b: %t", a != b)
	fmt.Printf("\na < b: %t", a < b)
	fmt.Printf("\na > b: %t", a > b)
	fmt.Printf("\na <= b: %t", a <= b)
	fmt.Printf("\na >= b: %t", a >= b)
}
