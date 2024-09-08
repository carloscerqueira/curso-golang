/*
1. Escreva um programa que imprima seu nome, idade e cidade natal usando
fmt.Println e fmt.Printf para formatar a saída.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite seu nome:")
	nome, _ := reader.ReadString('\n')

	fmt.Println("Digite sua idade:")
	idade, _ := reader.ReadString('\n')

	fmt.Println("Digite sua cidade natal:")
	cidade, _ := reader.ReadString('\n')

	fmt.Printf("Nome: %s", nome)
	fmt.Printf("Idade: %s", idade)
	fmt.Printf("Cidade natal: %s", cidade)
}
