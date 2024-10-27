/* Exercício 1: Escreva testes unitários para uma função que
calcula a média de uma lista de números inteiros. */

package main

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encotnrado foi %v."

func TestMedia(t *testing.T) {
	resultado := Media(4, 5, 6)
	esperado := 5

	if resultado != esperado {
		t.Errorf(erroPadrao, esperado, resultado)
	}
}
