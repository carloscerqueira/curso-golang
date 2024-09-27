/*
	Exercício 3: Crie uma função dividir que trate erros

de divisão por zero e teste com diferentes entradas.
*/
package main

import "fmt"

func divisao(x float64, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, fmt.Errorf("Não é possível realizar divisão por %.f", y)
	}
	return x / y, nil
}

func main() {
	resultado, err := divisao(7.8, 9.6)
	if err == nil {
		fmt.Printf("%.2f / %.2f = %.2f\n", 7.8, 9.6, resultado)
	} else {
		fmt.Println(err)
	}

	resultado, err = divisao(10.0, 0.00)
	if err == nil {
		fmt.Printf("%.2f / %.2f = %.2f\n", 10.0, 0.00, resultado)
	} else {
		fmt.Println(err)
	}

	resultado, err = divisao(10.0, 0.0)
	if err == nil {
		fmt.Printf("%.2f / %.2f = %.2f\n", 10.0, 0.0, resultado)
	} else {
		fmt.Println(err)
	}

	resultado, err = divisao(0.0, 10.0)
	if err == nil {
		fmt.Printf("%.2f / %.2f = %.2f\n", 0.0, 10.0, resultado)
	} else {
		fmt.Println(err)
	}

	resultado, err = divisao(6.0, 2.0)
	if err == nil {
		fmt.Printf("%.2f / %.2f = %.2f\n", 6.0, 2.0, resultado)
	} else {
		fmt.Println(err)
	}
}
