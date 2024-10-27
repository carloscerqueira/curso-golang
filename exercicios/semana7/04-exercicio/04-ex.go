package main

import (
	"fmt"
	"os"
)

func LerArquivo(arquivo string) (string, error) {
	body, err := os.ReadFile(arquivo)
	if err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo")
	}
	return string(body), nil
}
