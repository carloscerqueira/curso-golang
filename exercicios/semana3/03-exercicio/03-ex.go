/* Exercício 3: Crie um map para armazenar nomes e notas de alunos.
Em seguida, escreva uma função que imprima a média das notas. */

package main

import "fmt"

func main() {
	mapNotas := map[string][]float64{
		"José": {
			6.5,
			7.8,
			9.8,
			5.0,
		},
		"Maria": {
			6.0,
			10.0,
			8.0,
			6.0,
		},
		"João": {
			9.9,
			8.0,
			7.9,
			8.3,
		},
	}
	for nome, notas := range mapNotas {
		fmt.Printf("Nome: %s, Média: %.2f\n", nome, calculaMedia(notas))
	}
}

func calculaMedia(notas []float64) float64 {
	return calculaSoma(notas) / float64(len(notas))
}

func calculaSoma(notas []float64) float64 {
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	return soma
}
