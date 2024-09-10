package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// A variavel inicializada dessa forma, só existe dentro do escopo do bloco onde foi inicializada
	if i := numeroAleatorio(); i > 5 { // também suportado no switch
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
}
