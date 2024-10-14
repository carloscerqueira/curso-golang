/* Exercício 2: Crie um programa que simule um produtor e consumidor usando channels. */
package main

import "fmt"

func producer(mensagem string, ch chan string) {
	ch <- mensagem
}

func consumer(ch chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go producer("Olá", ch)
	go producer("Mundo", ch)
	consumer(ch)
	consumer(ch)
}
