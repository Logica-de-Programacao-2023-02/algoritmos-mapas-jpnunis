package main

import "fmt"

func pairFrequencies(numbers []int) map[[2]int]int {
	frequencyMap := make(map[[2]int]int)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			frequencyMap[pair]++
		}
	}

	return frequencyMap
}

func main() {
	numberSlice := []int{1, 2, 3, 1, 2, 4, 5, 3, 6}
	frequencyMap := pairFrequencies(numberSlice)

	fmt.Println("Pair Frequencies:")
	for pair, frequency := range frequencyMap {
		fmt.Printf("%v: %d\n", pair, frequency)
	}
}
