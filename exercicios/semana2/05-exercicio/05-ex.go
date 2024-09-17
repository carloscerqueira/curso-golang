/* 5. Crie um programa que peça ao usuário para adivinhar um número entre 1 e 10.
Use um laço for para permitir várias tentativas até que o usuário acerte. */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numero int
	for {
		fmt.Println("Tente adivinhar um número de 1 a 10")
		fmt.Scanln(&numero)
		if numero == (rand.Intn(10) + 1) {
			fmt.Println("Acertou!")
			break
		}
		fmt.Println("Errou!")
	}
}
