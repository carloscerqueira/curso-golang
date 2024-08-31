package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix

	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // x -= 1 ou x = x - 1
	fmt.Println(y)

	// Não podemos incrementar e decrementar dentro de uma função
	// é necessário fazer antes ou depois
	// fmt.Println(x == y--)

	// também não existe forma prefixada em go (--x ou ++y)
}
