/* Exercício 1: Crie uma struct que represente um livro,
com campos para título, autor e número de páginas.
Em seguida, crie um método para exibir as informações do livro. */

package main

import "fmt"

type Livro struct {
	titulo     string
	autor      string
	numPaginas int
}

func (l Livro) ExibirLivro() {
	fmt.Printf("Nome do livro: %s\nAutor: %s\nNúmero de páginas: %d\n", l.titulo, l.autor, l.numPaginas)
}

func main() {
	l := Livro{"Three ways to survive in a ruined world", "tls123", 314900}
	l.ExibirLivro()
}
