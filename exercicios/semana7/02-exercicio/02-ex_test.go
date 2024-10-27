/* Exercício 2: Crie uma função para calcular o fatorial de um número
e escreva um teste para verificar se a função trata corretamente valores negativos. */

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFatorial(t *testing.T) {
	casosDeTeste := []struct {
		nome     string
		numero   int
		esperado int
		err      error
	}{
		{"Sucesso", 5, 120, nil},
		{"Erro", -4, -1, fmt.Errorf("Número inválido: %d", -4)},
	}

	for _, tc := range casosDeTeste {
		t.Run(tc.nome, func(t *testing.T) {
			resultado, err := Fatorial(tc.numero)
			if tc.err != nil {
				assert.Error(t, tc.err, err)
			}
			if tc.err == nil {
				assert.Equal(t, tc.esperado, resultado)
			}
		})
	}
}
