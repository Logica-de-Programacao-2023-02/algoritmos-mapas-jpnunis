package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	return anagramMap
}

func sortString(s string) string {
	sortedRunes := []rune(s)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	words := []string{"listen", "silent", "enlist", "hello", "world", "dlrow"}
	anagramGroups := groupAnagrams(words)

	for key, group := range anagramGroups {
		fmt.Printf("Anagrams of %s: %v\n", group[0], group)
	}
}
