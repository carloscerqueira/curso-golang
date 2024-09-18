package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// funciona com slice, mas, não funciona com array
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)
}
