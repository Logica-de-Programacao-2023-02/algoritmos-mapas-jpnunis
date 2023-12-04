package main

import "fmt"

func characterFrequency(input string) map[rune]int {
	frequencyMap := make(map[rune]int)

	for _, char := range input {
		frequencyMap[char]++
	}

	return frequencyMap
}

func main() {
	inputString := "hello world"
	frequencyMap := characterFrequency(inputString)

	fmt.Println("Character frequency:")
	for char, frequency := range frequencyMap {
		fmt.Printf("%c: %d\n", char, frequency)
	}
}
