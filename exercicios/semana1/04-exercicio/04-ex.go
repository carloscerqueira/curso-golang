/* 4. Instancie dois números e realize as seguintes operações:
- soma
- subtração
- produto
- quociente
- resto da divisão (módulo) */

package main

import "fmt"

func main() {
	x, y := 42, 8
	fmt.Printf("Soma: %d", x+y)
	fmt.Printf("\nSubtração: %d", x-y)
	fmt.Printf("\nProduto: %d", x*y)
	fmt.Printf("\nQuociente: %d", x/y)
	fmt.Printf("\nResto da divisão: %d", x%y)
}
