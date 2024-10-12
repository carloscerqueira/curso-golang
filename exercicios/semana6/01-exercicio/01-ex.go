/*
	Exercício 1: Implemente um programa que lance múltiplas goroutines para calcular o

quadrado de números e use sync.WaitGroup para esperar que todas terminem.
*/
package main

import (
	"fmt"
	"sync"
)

func calculaQuadrado(numero int, wg *sync.WaitGroup) {
	fmt.Println(numero * numero)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go calculaQuadrado(4, &wg)
	wg.Add(1)
	go calculaQuadrado(6, &wg)
	wg.Wait()
}
