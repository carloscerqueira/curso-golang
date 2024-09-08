/* 5. Inicialize uma variável x com o valor 10 e use os operadores de atribuição (+=, -=, *=, /=, %=) para modificar
o valor dela, exibindo o resultado após cada operação. */

package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("Inicial: %d", x)
	x += 1
	fmt.Printf("\nx += 1: %d", x)
	x -= 1
	fmt.Printf("\nx -= 1: %d", x)
	x *= 2
	fmt.Printf("\nx *= 2: %d", x)
	x /= 2
	fmt.Printf("\nx /= 2: %d", x)
	x %= 3
	fmt.Printf("\nx %%= 3: %d", x)
}
