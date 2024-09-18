package main

import "fmt"

// A função init por convenção executa antes da função main
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
