package main

import "fmt"

func wordLetterCounts(input string) map[string]map[rune]int {
	wordCounts := make(map[string]map[rune]int)

	words := splitWords(input)

	for _, word := range words {
		wordCounts[word] = make(map[rune]int)
		for _, char := range word {
			wordCounts[word][char]++
		}
	}

	return wordCounts
}

func splitWords(input string) []string {
	var words []string
	currentWord := ""
	for _, char := range input {
		if char == ' ' || char == ',' || char == '.' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}
	return words
}

func main() {
	inputString := "hello world, golang is awesome."
	wordCounts := wordLetterCounts(inputString)

	fmt.Println("Word Letter Counts:")
	for word, letterCounts := range wordCounts {
		fmt.Printf("%s: %v\n", word, letterCounts)
	}
}
