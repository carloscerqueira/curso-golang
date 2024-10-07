/*
	Exercício 4: Implemente um sistema de cadastro de pessoas utilizando structs e métodos.

As pessoas podem ser estudantes ou professores,
e a informação a ser exibida deve ser diferente para cada tipo.
*/
package main

import "fmt"

type Pessoa interface {
	cadastraPessoa()
	imprimeDados()
}

type Estudante struct {
	nome      string
	idade     int
	interesse string
}

type Professor struct {
	nome    string
	idade   int
	materia string
}

func (e *Estudante) cadastraPessoa() {
	fmt.Println("Insira os dados do estudante\nNome: ")
	fmt.Scanln(&e.nome)
	fmt.Println("Idade: ")
	fmt.Scanln(&e.idade)
	fmt.Println("Interesse: ")
	fmt.Scanln(&e.interesse)
}

func (p *Professor) cadastraPessoa() {
	fmt.Println("Insira os dados do professor\nNome: ")
	fmt.Scanln(&p.nome)
	fmt.Println("Idade: ")
	fmt.Scanln(&p.idade)
	fmt.Println("Matéria: ")
	fmt.Scanln(&p.materia)
}

func (e *Estudante) imprimeDados() {
	fmt.Printf("Dados do estudante\nNome: %s, idade: %d, interesse: %s\n", e.nome, e.idade, e.interesse)
}

func (p *Professor) imprimeDados() {
	fmt.Printf("Dados do professor\nNome: %s, idade: %d, materia: %s\n", p.nome, p.idade, p.materia)
}

func main() {
	e := Estudante{}
	p := Professor{}
	e.cadastraPessoa()
	e.imprimeDados()
	p.cadastraPessoa()
	p.imprimeDados()
}
