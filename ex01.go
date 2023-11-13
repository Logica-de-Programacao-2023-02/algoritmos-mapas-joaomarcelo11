package main

import (
	"fmt"
	"strings"
)

func contarPalavras(texto string) map[string]int {
	palavras := strings.Fields(texto)
	ocorrencias := make(map[string]int)
	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}
	return ocorrencias
}

func main() {
	//Últimos 10 campeões brasileiros
	x := "Palmeiras , Atlético , Flamengo , Flamengo , Palmeiras , Corinthians , Palmeiras , Corinthians ," +
		" Cruzeiro , Cruzeiro"
	fmt.Println(contarPalavras(x))
}