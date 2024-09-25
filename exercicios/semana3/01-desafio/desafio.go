/* Desafio: Delete um elemento da slice utilizando as operações aqui ensinadas */

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	/* como o append é uma função variática, ao concatenar slices, preciso colocar "..."
	no segundo parametro */
	slice = append(slice[:3], slice[3+1:]...)
	fmt.Println(slice)
}
