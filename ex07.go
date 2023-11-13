package main

import (
	"fmt"
	"strings"
)

func countLetters(word string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, letter := range word {
		letterCount[letter]++
	}

	return letterCount
}

func wordLetterCount(sentence string) map[string]map[rune]int {
	wordCount := make(map[string]map[rune]int)

	words := strings.Fields(sentence)

	for _, word := range words {
		wordCount[word] = countLetters(word)
	}

	return wordCount
}

func main() {
	sentence := "Hello, world!"

	wordLetterMap := wordLetterCount(sentence)

	for word, letterCount := range wordLetterMap {
		fmt.Printf("%s: %v\n", word, letterCount)
	}
}