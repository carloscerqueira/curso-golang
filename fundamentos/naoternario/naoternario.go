package main

import "fmt"

// Golang não tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	// NÃO EXISTE
	// return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main() {
	resultado := obterResultado(6.2)

	fmt.Println("O aluno foi", resultado)
}
