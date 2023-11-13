package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func groupAnagrams(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		sorted := sortString(strings.ToLower(word))
		anagramGroups[sorted] = append(anagramGroups[sorted], word)
	}

	return anagramGroups
}

func main() {
	wordSlice := []string{"cat", "dog", "tac", "god", "act", "good"}

	anagramMap := groupAnagrams(wordSlice)

	for _, group := range anagramMap {
		fmt.Println(group)
	}
}