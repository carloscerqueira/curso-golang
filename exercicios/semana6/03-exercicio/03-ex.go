/*
	Exercício 3: Implemente um sistema que gere números aleatórios em uma goroutine e utilize

outra goroutine para somar esses números recebidos via channels.
*/
package main

import (
	"fmt"
	"math/rand"
)

func geraAleatorios(c chan int) {
	for i := 0; i < 10; i++ {
		c <- rand.Intn(10)
	}
	close(c)
}

func main() {
	c := make(chan int)
	go geraAleatorios(c)
	soma := 0
	for numero := range c {
		soma += numero
	}
	fmt.Println(soma)
}
