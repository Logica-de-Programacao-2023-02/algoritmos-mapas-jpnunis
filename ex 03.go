package main

import "fmt"

func somarValoresDoMapa(mapaValores map[string]int) int {
	soma := 0

	for _, valor := range mapaValores {
		soma += valor
	}

	return soma
}

func main() {
	valores := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	resultado := somarValoresDoMapa(valores)

	fmt.Printf("A soma dos valores do mapa é: %d\n", resultado)
}
