package main

import (
	"fmt"
	"time"

	"github.com/carloscerqueira/gethtmltitle"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := gethtmltitle.Title(url1)
	c2 := gethtmltitle.Title(url2)
	c3 := gethtmltitle.Title(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.wikipedia.com",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
