package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("linha")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516

	xs := fmt.Sprint(x)

	// Só da pra concatenar se a variavel for uma string
	fmt.Println("O valor de x é: " + xs)
	fmt.Println("O valor de x é:", x)

	// Com Printf fica melhor
	// \n pula linha já que o Printf não pula
	fmt.Printf("O valor de x é: %f\n", x)
	fmt.Printf("O valor de x é: %.2f\n", x)

	a := 1
	b := 1.9999
	c := false
	d := "aoba"
	fmt.Printf("%d %f %.1f %t %s", a, b, b, c, d)
	// Também é possível utilizar %v
	/* %v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names */

	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
