/* Exercício 4: Implemente um programa em Go que simule o funcionamento de um banco, onde múltiplas
goroutines podem realizar saques e depósitos simultaneamente em uma conta compartilhada.
Utilize um sync.Mutex para garantir que as operações de saque e depósito sejam feitas de
forma segura e sem condições de corrida. */

package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var saldo float64

func depositar(wg *sync.WaitGroup, valorDeposito float64) {
	defer wg.Done()
	mu.Lock()
	saldo += valorDeposito
	mu.Unlock()
}

func sacar(wg *sync.WaitGroup, valorSaque float64) {
	defer wg.Done()
	mu.Lock()
	saldo -= valorSaque
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go depositar(&wg, 34.65)
	wg.Add(1)
	go depositar(&wg, 34.65)
	wg.Add(1)
	go depositar(&wg, 34.65)
	wg.Add(1)
	go sacar(&wg, 12.90)
	wg.Wait()
	fmt.Printf("Saldo final: %.2f", saldo)
}
