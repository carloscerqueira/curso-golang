/* Desafio: Demonstre que é possível usar o for como While presente em outras linguagens */
package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
