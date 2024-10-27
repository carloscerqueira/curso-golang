package main

func IsPalindromo(palavra string) bool {
	caracteres := []rune(palavra)
	for i, j := 0, len(caracteres)-1; i < j; i, j = i+1, j-1 {
		caracteres[i], caracteres[j] = caracteres[j], caracteres[i]
	}
	palavraReversa := string(caracteres)
	return palavraReversa == palavra
}
