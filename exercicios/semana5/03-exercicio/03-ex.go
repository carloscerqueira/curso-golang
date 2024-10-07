/* Exercício 3: Crie uma função que receba uma lista de valores do tipo interface{}
e imprima o tipo de cada valor usando type assertion. */

package main

import "fmt"

func imprimeTipoDado(valores ...interface{}) {
	for _, v := range valores {
		valorString, isString := v.(string)
		valorInteiro, isInteiro := v.(int)
		valorReal, isReal := v.(float64)
		valorBooleano, isBooleano := v.(bool)
		switch {
		case isString:
			fmt.Println("É uma string:", valorString)
		case isInteiro:
			fmt.Println("É um inteiro:", valorInteiro)
		case isReal:
			fmt.Println("É um real:", valorReal)
		case isBooleano:
			fmt.Println("É um booleano:", valorBooleano)
		default:
			fmt.Println("Tipo não mapeado")
		}
	}
}

func main() {
	imprimeTipoDado("String", 42, 42.50, false)
}
