/* 2. Declare variáveis de diferentes tipos: int, float64, bool e string. Atribua valores a elas e imprima-os no console.
Depois, declare uma variável sem inicializá-la e verifique o valor "zero" padrão que o Go atribui a ela. */

package main

import "fmt"

func main() {
	a := 42
	b := 42.42
	c := true
	d := "Olá!"

	fmt.Printf(`
	Variáveis inicializadas
	int: %d
	float64: %f
	bool: %t
	string: %s`, a, b, c, d)

	var inteiro int
	var flutuante64 float64
	var booleana bool
	var texto string

	fmt.Printf(`
	Valores "zero" das variáveis
	int: %d
	float64: %f
	bool: %t
	string: %s`, inteiro, flutuante64, booleana, texto)

}
