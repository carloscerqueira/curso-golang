package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 200)
				break
			}
		}
	}
	close(c) // indica que não vai mais haver inserção no canal, evita deadlock
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
