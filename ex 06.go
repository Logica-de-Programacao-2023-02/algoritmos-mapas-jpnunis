package main

import "fmt"

func mergeWordCounts(wordCountsList []map[string]int) map[string]int {
	mergedWordCounts := make(map[string]int)

	for _, wordCounts := range wordCountsList {
		for word, count := range wordCounts {
			mergedWordCounts[word] += count
		}
	}

	return mergedWordCounts
}

func main() {
	wordCountsList := []map[string]int{
		{"hello": 2, "world": 1, "go": 3},
		{"hello": 1, "world": 2, "golang": 4},
		{"go": 2, "gopher": 5},
	}

	mergedCounts := mergeWordCounts(wordCountsList)

	fmt.Println("Merged Word Counts:")
	for word, count := range mergedCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
