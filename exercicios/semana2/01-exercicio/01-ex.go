/* 1. Crie um programa que receba a idade de uma pessoa e diga se ela é maior ou menor de idade, usando if-else. */
package main

import (
	"fmt"
)

func main() {

	var idade int
	fmt.Println("Digite sua idade")
	fmt.Scan(&idade)

	if idade >= 18 {
		fmt.Println("Maior de idade!")
	} else {
		fmt.Println("Menor de idade!")
	}
}
