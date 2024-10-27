package main

import "fmt"

func Fatorial(num int) (int, error) {
	if num < 0 {
		return -1, fmt.Errorf("Número inválido: %d", num)
	} else if num == 1 {
		return 1, nil
	} else {
		fatorialAnterior, _ := Fatorial(num - 1)
		return fatorialAnterior * num, nil
	}
}
