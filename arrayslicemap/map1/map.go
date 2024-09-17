package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98771235412] = "Pedro"
	aprovados[85719120394] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[85719120394])
	delete(aprovados, 85719120394)
	fmt.Println(aprovados[85719120394])
	fmt.Println(aprovados)
}
