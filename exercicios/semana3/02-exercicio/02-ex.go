/* Exercício 2: Crie um slice e adicione elementos a ele
dinamicamente até seu tamanho ultrapassar 10. */

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Tamanho inicial do slice: %d\n", len(slice))
	slice = append(slice, 10, 11, 12, 13, 14, 15)
	fmt.Printf("Tamanho do slice após append: %d\n", len(slice))
}
