/* Exercício 3: Escreva testes para uma função que verifica se uma string é um palíndromo. */

package main

import "testing"

func TestIsPalindromo(t *testing.T) {
	casosDeTeste := []struct {
		nome     string
		palavra  string
		esperado bool
	}{
		{"Palindromo", "arara", true},
		{"Não palindromo", "rato", false},
	}

	for _, tc := range casosDeTeste {
		t.Run(tc.nome, func(t *testing.T) {
			resultado := IsPalindromo(tc.palavra)
			if resultado != tc.esperado {
				t.Errorf("esperado: %v, encontrado: %v", tc.esperado, resultado)
			}
		})
	}
}
