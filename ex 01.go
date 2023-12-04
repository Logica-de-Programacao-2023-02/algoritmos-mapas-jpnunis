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
	textoExemplo := "go é uma linguagem de programação go é poderosa e eficiente go é divertida de aprender"

	resultado := contarPalavras(textoExemplo)

	fmt.Println("Ocorrência de cada palavra:")
	for palavra, ocorrencias := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencias)
	}
}
