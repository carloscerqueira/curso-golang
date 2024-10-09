package main

import "fmt"

type ByteSize int

const (
	_           = iota             // ignore first value (0) by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // bitshift to the left
	MB
	GB
	TB
	PB
	EB
	// ZB
	// YB
)

func main() {
	fmt.Printf("KB: %d \t\t\t\t\t %b\n", KB, KB)
	fmt.Printf("MB: %d \t\t\t\t %b\n", MB, MB)
	fmt.Printf("GB: %d \t\t\t\t %b\n", GB, GB)
	fmt.Printf("TB: %d \t\t\t %b\n", TB, TB)
	fmt.Printf("PB: %d \t\t %b\n", PB, PB)
	fmt.Printf("EB: %d \t %b\n", EB, EB)
}

/*
PB		    1125899906842624
EB		 1152921504606846976
int		18446744073709551615
*/
