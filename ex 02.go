package main

import "fmt"

func mesclarMapas(mapa1, mapa2 map[string]int) map[string]int {
	resultado := make(map[string]int)

	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {
	mapa1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	mapa2 := map[string]int{
		"b": 4,
		"c": 5,
		"d": 6,
	}

	resultado := mesclarMapas(mapa1, mapa2)

	fmt.Println("Mapa resultante da fusão:")
	for chave, valor := range resultado {
		fmt.Printf("%s: %d\n", chave, valor)
	}
}
