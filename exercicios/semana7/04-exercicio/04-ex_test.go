/* Exercício 4: Implemente uma função que lê e retorna o conteúdo de um arquivo.
Crie um teste que simula a leitura de arquivos inválidos e verifique se a função
lida corretamente com erros. */

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLerArquivo(t *testing.T) {
	casosDeTeste := []struct {
		nome     string
		arquivo  string
		esperado string
		err      error
	}{
		{"Sucesso", "arquivo.txt", "hello go", nil},
		{"Erro", "naoExiste.txt", "", fmt.Errorf("erro ao ler o arquivo")},
	}

	for _, tc := range casosDeTeste {
		t.Run(tc.nome, func(t *testing.T) {
			resultado, err := LerArquivo(tc.arquivo)
			if tc.err != nil {
				assert.Error(t, tc.err, err)
			}
			if tc.err == nil {
				assert.Equal(t, tc.esperado, resultado)
			}
		})
	}
}
