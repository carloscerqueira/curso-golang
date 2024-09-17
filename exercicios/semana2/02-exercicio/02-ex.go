/*
 2. Crie um programa que simule um semáforo usando switch.

Escaneie o valor e exiba mensagens para os estados
"Vermelho", "Amarelo" e "Verde". (Dica, usar a função fmt.Scanln())
*/
package main

import "fmt"

func main() {
	var estado string
	fmt.Println("Digite o valor do semáforo")
	fmt.Scanln(&estado)
	switch estado {
	case "Vermelho":
		fmt.Println("Pare")
	case "Amarelo":
		fmt.Println("Atenção")
	case "Verde":
		fmt.Println("Siga em frente")
	default:
		fmt.Println("Valor inválido")
	}
}
