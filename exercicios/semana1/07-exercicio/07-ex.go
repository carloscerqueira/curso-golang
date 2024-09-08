/* 7. Escreva um programa que incialize três variáveis booleanas (true ou false) e exiba os resultados das operações lógicas
combinadas (&&, ||, !). */

package main

import "fmt"

func main() {
	a, b, c := true, false, true
	fmt.Printf("a = %t\nb = %t\nc = %t", a, b, c)
	fmt.Printf("\na && b: %t", a && b)
	fmt.Printf("\na && c: %t", a && c)
	fmt.Printf("\nb && c: %t", b && c)
	fmt.Printf("\na || b: %t", a || b)
	fmt.Printf("\na || c: %t", a || c)
	fmt.Printf("\nb || c: %t", b || c)
	fmt.Printf("\n!a: %t", !a)
	fmt.Printf("\n!b: %t", !b)
	fmt.Printf("\n!c: %t", !c)
}
