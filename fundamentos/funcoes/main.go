package main

/*
	Dependendo do valor da variavel GO111MODULE a partir da versão 1.11 do golang,
	é preciso iniciar um modulo para poder rodar com a main
	em um arquivo separado, mesmo que esteja no mesmo package
*/

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}
