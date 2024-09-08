/* 1. Escreva um programa que imprima seu nome, idade e cidade natal usando
fmt.Println e fmt.Printf para formatar a saída. */

package main

import "fmt"

func main() {
	nome := "Fulano de Tal"
	idade := 40
	cidade := "São Paulo"
	fmt.Println("Nome:", nome)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Cidade natal: %s\n", cidade)
}
